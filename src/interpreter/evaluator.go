package interpreter

import (
	. "byx-script-go/src/common"
	. "byx-script-go/src/parser"
	"container/list"
	"fmt"
	"reflect"
	"strings"
)

// 表达式求值
func evaluate(node any, scope Scope) Value {
	switch node.(type) {
	// 字面量（整数、浮点数、字符串、布尔值）
	case Literal:
		return node.(Literal).V
	// 变量引用
	case Var:
		return scope.GetVar(node.(Var).VarName)
	// 列表字面量
	case ListLiteral:
		return evalListLiteral(node.(ListLiteral), scope)
	// 函数字面量
	case CallableLiteral:
		return evalCallableLiteral(node.(CallableLiteral), scope)
	// 对象字面量
	case ObjectLiteral:
		return evalObjectLiteral(node.(ObjectLiteral), scope)
	// 一元表达式
	case UnaryExpr:
		return evalUnaryExpr(node.(UnaryExpr), scope)
	// 二元表达式
	case BinaryExpr:
		return evalBinaryExpr(node.(BinaryExpr), scope)
	// 函数调用
	case Call:
		return evalCall(node.(Call), scope)
	// 字段访问
	case FieldAccess:
		return evalFieldAccess(node.(FieldAccess), scope)
	// 下标访问
	case Subscript:
		return evalSubscript(node.(Subscript), scope)
	// 未知表达式类型
	default:
		panic(fmt.Sprintf("unknown expression type: %T", node))
	}
}

func getListIndex(lst *list.List, index int) Value {
	for e := lst.Front(); e != nil; e = e.Next() {
		if index == 0 {
			return e.Value.(Value)
		}
		index--
	}
	return UndefinedValue()
}

func evalSubscript(node Subscript, scope Scope) Value {
	e := evaluate(node.Expr, scope)
	sub := evaluate(node.Sub, scope)
	if !sub.IsInteger() {
		panic(fmt.Sprintf("subscript must be integer: %v", sub))
	}
	index := sub.GetInteger()
	if e.IsString() {
		return StringValue(string([]rune(e.GetString())[index]))
	} else if e.IsList() {
		return getListIndex(e.GetList(), index)
	}
	return UndefinedValue()
}

func evalObjectField(fields Fields, name string) Value {
	v, exist := fields[name]
	if exist {
		return v
	}
	return UndefinedValue()
}

func evalListField(lst *list.List, name string) Value {
	switch name {
	case "length":
		return ListLength(lst)
	case "addFirst":
		return ListAddFirst(lst)
	case "removeFirst":
		return ListRemoveFirst(lst)
	case "addLast":
		return ListAddLast(lst)
	case "removeLast":
		return ListRemoveLast(lst)
	case "insert":
		return ListInsert(lst)
	case "remove":
		return ListRemove(lst)
	case "copy":
		return ListCopy(lst)
	case "isEmpty":
		return ListIsEmpty(lst)
	default:
		return UndefinedValue()
	}
}

func evalStringField(s string, name string) Value {
	switch name {
	case "length":
		return StringLength(s)
	case "substring":
		return StringSubstring(s)
	case "concat":
		return StringConcat(s)
	case "charAt":
		return StringCharAt(s)
	case "codeAt":
		return StringCodeAt(s)
	case "toUpper":
		return StringToUpper(s)
	case "toLower":
		return StringToLower(s)
	default:
		return UndefinedValue()
	}
}

func evalFieldAccess(node FieldAccess, scope Scope) Value {
	v := evaluate(node.Expr, scope)
	name := node.Name
	if v.IsString() {
		return evalStringField(v.GetString(), name)
	} else if v.IsList() {
		return evalListField(v.GetList(), name)
	} else if v.IsObject() {
		return evalObjectField(v.GetFields(), name)
	}
	return UndefinedValue()
}

func evalCall(node Call, scope Scope) Value {
	callee := evaluate(node.Callee, scope)
	if callee.IsCallable() {
		var args []Value
		for _, a := range node.Args {
			args = append(args, evaluate(a, scope))
		}
		return callee.GetCallable()(args)
	}
	panic(fmt.Sprintf("%v is not callable", callee))
}

func cannotApplyBinaryOp(op string, v1 Value, v2 Value) Value {
	panic(fmt.Sprintf("operator %s cannot apply between %s and %s", op, v1.GetTypeId(), v2.GetTypeId()))
}

func evalAdd(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsString() || v2.IsString() {
		return StringValue(v1.String() + v2.String())
	}

	if v1.IsInteger() && v2.IsInteger() {
		return IntegerValue(v1.GetInteger() + v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return DoubleValue(float64(v1.GetInteger()) + v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return DoubleValue(v1.GetDouble() + float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return DoubleValue(v1.GetDouble() + v2.GetDouble())
	} else if v1.IsObject() {
		fields := v1.GetFields()
		if v, exist := fields["_add"]; exist && v.IsCallable() {
			return v.GetCallable()([]Value{v2})
		}
	}

	return cannotApplyBinaryOp("+", v1, v2)
}

func evalSub(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsInteger() && v2.IsInteger() {
		return IntegerValue(v1.GetInteger() - v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return DoubleValue(float64(v1.GetInteger()) - v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return DoubleValue(v1.GetDouble() - float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return DoubleValue(v1.GetDouble() - v2.GetDouble())
	} else if v1.IsObject() {
		fields := v1.GetFields()
		if v, exist := fields["_sub"]; exist && v.IsCallable() {
			return v.GetCallable()([]Value{v2})
		}
	}

	return cannotApplyBinaryOp("-", v1, v2)
}

func evalMul(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsInteger() && v2.IsInteger() {
		return IntegerValue(v1.GetInteger() * v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return DoubleValue(float64(v1.GetInteger()) * v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return DoubleValue(v1.GetDouble() * float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return DoubleValue(v1.GetDouble() * v2.GetDouble())
	} else if v1.IsObject() {
		fields := v1.GetFields()
		if v, exist := fields["_mul"]; exist && v.IsCallable() {
			return v.GetCallable()([]Value{v2})
		}
	}

	return cannotApplyBinaryOp("*", v1, v2)
}

func evalDiv(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsInteger() && v2.IsInteger() {
		return IntegerValue(v1.GetInteger() / v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return DoubleValue(float64(v1.GetInteger()) / v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return DoubleValue(v1.GetDouble() / float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return DoubleValue(v1.GetDouble() / v2.GetDouble())
	} else if v1.IsObject() {
		fields := v1.GetFields()
		if v, exist := fields["_div"]; exist && v.IsCallable() {
			return v.GetCallable()([]Value{v2})
		}
	}

	return cannotApplyBinaryOp("/", v1, v2)
}

func evalRem(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsInteger() && v2.IsInteger() {
		return IntegerValue(v1.GetInteger() % v2.GetInteger())
	}

	return cannotApplyBinaryOp("%", v1, v2)
}

func evalLessThan(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsString() && v2.IsString() {
		return BoolValue(strings.Compare(v1.GetString(), v2.GetString()) < 0)
	} else if v1.IsInteger() && v2.IsInteger() {
		return BoolValue(v1.GetInteger() < v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return BoolValue(float64(v1.GetInteger()) < v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return BoolValue(v1.GetDouble() < float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return BoolValue(v1.GetDouble() < v2.GetDouble())
	}

	return cannotApplyBinaryOp("<", v1, v2)
}

func evalLessEqualThan(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsString() && v2.IsString() {
		return BoolValue(strings.Compare(v1.GetString(), v2.GetString()) <= 0)
	} else if v1.IsInteger() && v2.IsInteger() {
		return BoolValue(v1.GetInteger() <= v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return BoolValue(float64(v1.GetInteger()) <= v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return BoolValue(v1.GetDouble() <= float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return BoolValue(v1.GetDouble() <= v2.GetDouble())
	}

	return cannotApplyBinaryOp("<=", v1, v2)
}

func evalGreaterThan(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsString() && v2.IsString() {
		return BoolValue(strings.Compare(v1.GetString(), v2.GetString()) > 0)
	} else if v1.IsInteger() && v2.IsInteger() {
		return BoolValue(v1.GetInteger() > v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return BoolValue(float64(v1.GetInteger()) > v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return BoolValue(v1.GetDouble() > float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return BoolValue(v1.GetDouble() > v2.GetDouble())
	}

	return cannotApplyBinaryOp(">", v1, v2)
}

func evalGreaterEqualThan(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsString() && v2.IsString() {
		return BoolValue(strings.Compare(v1.GetString(), v2.GetString()) >= 0)
	} else if v1.IsInteger() && v2.IsInteger() {
		return BoolValue(v1.GetInteger() >= v2.GetInteger())
	} else if v1.IsInteger() && v2.IsDouble() {
		return BoolValue(float64(v1.GetInteger()) >= v2.GetDouble())
	} else if v1.IsDouble() && v2.IsInteger() {
		return BoolValue(v1.GetDouble() >= float64(v2.GetInteger()))
	} else if v1.IsDouble() && v2.IsDouble() {
		return BoolValue(v1.GetDouble() >= v2.GetDouble())
	}

	return cannotApplyBinaryOp(">=", v1, v2)
}

func evalEqual(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	v2 := evaluate(rhs, scope)

	if v1.IsObject() && v2.IsObject() {
		return BoolValue(fmt.Sprintf("%p", v1.GetData()) == fmt.Sprintf("%p", v2.GetData()))
	} else {
		return BoolValue(reflect.DeepEqual(v1.GetData(), v2.GetData()))
	}
}

func evalNotEqual(lhs any, rhs any, scope Scope) Value {
	return BoolValue(!evalEqual(lhs, rhs, scope).GetBool())
}

func evalAnd(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	if v1.IsBool() && !v1.GetBool() {
		return BoolValue(false)
	}

	v2 := evaluate(rhs, scope)
	if v2.IsBool() {
		return BoolValue(v2.GetBool())
	}

	return cannotApplyBinaryOp("&&", v1, v2)
}

func evalOr(lhs any, rhs any, scope Scope) Value {
	v1 := evaluate(lhs, scope)
	if v1.IsBool() && v1.GetBool() {
		return BoolValue(true)
	}
	v2 := evaluate(rhs, scope)
	if v2.IsBool() {
		return BoolValue(v2.GetBool())
	}

	return cannotApplyBinaryOp("||", v1, v2)
}

func evalBinaryExpr(node BinaryExpr, scope Scope) Value {
	op := node.Op
	lhs := node.Lhs
	rhs := node.Rhs
	switch op {
	case "+":
		return evalAdd(lhs, rhs, scope)
	case "-":
		return evalSub(lhs, rhs, scope)
	case "*":
		return evalMul(lhs, rhs, scope)
	case "/":
		return evalDiv(lhs, rhs, scope)
	case "%":
		return evalRem(lhs, rhs, scope)
	case "<":
		return evalLessThan(lhs, rhs, scope)
	case "<=":
		return evalLessEqualThan(lhs, rhs, scope)
	case ">":
		return evalGreaterThan(lhs, rhs, scope)
	case ">=":
		return evalGreaterEqualThan(lhs, rhs, scope)
	case "==":
		return evalEqual(lhs, rhs, scope)
	case "!=":
		return evalNotEqual(lhs, rhs, scope)
	case "&&":
		return evalAnd(lhs, rhs, scope)
	case "||":
		return evalOr(lhs, rhs, scope)
	default:
		panic(fmt.Sprintf("unknown binary operator: %v", op))
	}
}

func evalNot(e any, scope Scope) Value {
	v := evaluate(e, scope)
	if v.IsBool() {
		return BoolValue(!v.GetBool())
	} else {
		panic(fmt.Sprintf("operator ! cannot apply on: %v", v))
	}
}

func evalNeg(e any, scope Scope) Value {
	v := evaluate(e, scope)
	if v.IsInteger() {
		return IntegerValue(-v.GetInteger())
	} else if v.IsDouble() {
		return DoubleValue(-v.GetDouble())
	} else {
		panic(fmt.Sprintf("operator - cannot apply on: %v", v))
	}
}

func evalUnaryExpr(node UnaryExpr, scope Scope) Value {
	op := node.Op
	e := node.Expr
	switch op {
	case "-":
		return evalNeg(e, scope)
	case "!":
		return evalNot(e, scope)
	default:
		panic(fmt.Sprintf("unknown unary operator: %v", op))
	}
}

func evalObjectLiteral(node ObjectLiteral, scope Scope) Value {
	fields := Fields{}
	for k, v := range node.Fields {
		fields[k] = evaluate(v, scope)
	}
	return ObjectValue(fields)
}

func evalCallableLiteral(node CallableLiteral, scope Scope) Value {
	return CallableValue(func(args []Value) (retVal Value) {
		// 传递实参
		newScope := NewScope(scope)
		for i, p := range node.Params {
			if i < len(args) {
				newScope.DeclareVar(p, args[i])
			} else {
				newScope.DeclareVar(p, UndefinedValue())
			}
		}

		// 捕获函数返回异常
		defer func() {
			if r := recover(); r != nil {
				if e, ok := r.(ReturnException); ok {
					retVal = e.retVal
				} else {
					panic(r)
				}
			}
		}()

		// 执行函数体
		execute(node.Body, newScope)
		return UndefinedValue()
	})
}

func evalListLiteral(node ListLiteral, scope Scope) Value {
	elems := list.New()
	for _, n := range node.Elems {
		elems.PushBack(evaluate(n, scope))
	}
	return ListValue(elems)
}
