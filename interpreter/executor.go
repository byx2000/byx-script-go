package interpreter

import (
	. "byx-script-go/common"
	. "byx-script-go/parser"
	"container/list"
	"fmt"
)

type BreakException struct {
}

type ContinueException struct {
}

type ReturnException struct {
	retVal Value
}

type ThrowException struct {
	value Value
}

const (
	NORMAL = iota
	BREAK
	CONTINUE
)

// 执行语句
func execute(node any, scope Scope) {
	switch node.(type) {
	// 变量声明
	case VarDeclare:
		execVarDeclare(node.(VarDeclare), scope)
	// 赋值
	case Assign:
		execAssign(node.(Assign), scope)
	// 语句块
	case Block:
		execBlock(node.(Block), scope)
	// if语句
	case If:
		execIf(node.(If), scope)
	// while语句
	case While:
		execWhile(node.(While), scope)
	// for语句
	case For:
		execFor(node.(For), scope)
	// try语句
	case Try:
		execTry(node.(Try), scope)
	// throw语句
	case Throw:
		value := evaluate(node.(Throw).Expr, scope)
		panic(ThrowException{value})
	// 跳出循环
	case Break:
		panic(BreakException{})
	// 跳过当前循环
	case Continue:
		panic(ContinueException{})
	// 函数返回
	case Return:
		retVal := evaluate(node.(Return).RetVal, scope)
		panic(ReturnException{retVal})
	// 表达式语句
	case ExprStatement:
		evaluate(node.(ExprStatement).Expr, scope)
	// 程序
	case Program:
		for _, s := range node.(Program).Stmts {
			execute(s, scope)
		}
	// 未知语句类型
	default:
		panic(fmt.Sprintf("unknown statement type: %T", node))
	}
}

func execTry(node Try, scope Scope) {
	defer func() {
		if e := recover(); e != nil {
			if t, ok := e.(ThrowException); ok {
				newScope := NewScope(scope)
				newScope.DeclareVar(node.CatchVar, t.value)
				execute(node.CatchBranch, newScope)
			} else {
				execute(node.FinallyBranch, scope)
				panic(e)
			}
		}
		execute(node.FinallyBranch, scope)
	}()

	execute(node.TryBranch, scope)
}

func execFor(node For, scope Scope) {
	scope = NewScope(scope)
	for execute(node.Init, scope); getCondition(evaluate(node.Cond, scope)); execute(node.Update, scope) {
		r := doLoop(node.Body, scope)
		if r == BREAK {
			break
		}
	}
}

func doLoop(body any, scope Scope) (ret int) {
	defer func() {
		if e := recover(); e != nil {
			switch e.(type) {
			case BreakException:
				ret = BREAK
			case ContinueException:
				ret = CONTINUE
			default:
				panic(e)
			}
		}
	}()

	execute(body, scope)
	return NORMAL
}

func getCondition(cond Value) bool {
	if cond.IsBool() {
		return cond.GetBool()
	}
	panic(fmt.Sprintf("condition of if, while, for statement must be bool value: %s", cond.String()))
}

func execWhile(node While, scope Scope) {
	for {
		if !getCondition(evaluate(node.Cond, scope)) {
			break
		}
		r := doLoop(node.Body, scope)
		if r == BREAK {
			break
		}
	}
}

func execIf(node If, scope Scope) {
	for _, c := range node.Cases {
		cond := evaluate(c.First, scope)
		if getCondition(cond) {
			execute(c.Second, scope)
			return
		}
	}
	execute(node.ElseBranch, scope)
}

func execBlock(node Block, scope Scope) {
	newScope := NewScope(scope)
	for _, s := range node.Stmts {
		execute(s, newScope)
	}
}

func setListIndex(lst *list.List, index int, v Value) {
	for e := lst.Front(); e != nil; e = e.Next() {
		if index == 0 {
			e.Value = v
			return
		}
		index--
	}
}

func execAssign(node Assign, scope Scope) {
	lhs := node.Lhs
	rhs := node.Rhs
	switch lhs.(type) {
	// 变量赋值
	case Var:
		scope.SetVar(lhs.(Var).VarName, evaluate(rhs, scope))
	// 字段赋值
	case FieldAccess:
		e := evaluate(lhs.(FieldAccess).Expr, scope)
		if e.IsObject() {
			fields := e.GetFields()
			name := lhs.(FieldAccess).Name
			fields[name] = evaluate(rhs, scope)
			return
		}
		panic(fmt.Sprintf("%v is not object", e))
	// 下标赋值
	case Subscript:
		e := evaluate(lhs.(Subscript).Expr, scope)
		if e.IsList() {
			sub := evaluate(lhs.(Subscript).Sub, scope)
			if sub.IsInteger() {
				setListIndex(e.GetList(), sub.GetInteger(), evaluate(rhs, scope))
				return
			}
			panic(fmt.Sprintf("%v must be integer", sub))
		}
		panic(fmt.Sprintf("%v is not list", e))
	default:
		panic(fmt.Sprintf("%v is not assignable", lhs))
	}
}

func execVarDeclare(node VarDeclare, scope Scope) {
	varName := node.VarName
	value := evaluate(node.Value, scope)
	scope.DeclareVar(varName, value)
}
