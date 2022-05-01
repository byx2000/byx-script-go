package parser

import (
	. "byx-script-go/src/common"
	. "byx-script-go/src/parserc"
	"fmt"
	"strconv"
)

func toString(v any) string {
	switch v.(type) {
	case rune:
		return fmt.Sprintf("%c", v)
	default:
		return fmt.Sprintf("%s", v)
	}
}

func join(list any) any {
	str := ""
	for _, e := range list.([]interface{}) {
		str += toString(e)
	}
	return str
}

func toInt(s any) int {
	v, _ := strconv.Atoi(s.(string))
	return v
}

func toFloat(s any) float64 {
	v, _ := strconv.ParseFloat(s.(string), 64)
	return v
}

func toBool(s any) bool {
	r, _ := strconv.ParseBool(s.(string))
	return r
}

func buildBinaryExpr(p any) any {
	e := p.(Pair).First
	for _, pp := range p.(Pair).Second.([]any) {
		op := pp.(Pair).First.(string)
		e = BinaryExpr{op, e, pp.(Pair).Second}
	}
	return e
}

func buildExprElem(p any) any {
	e := p.(Pair).First
	for _, o := range p.(Pair).Second.([]any) {
		switch o.(type) {
		case []any:
			e = Call{e, o.([]any)}
		case string:
			e = FieldAccess{e, o.(string)}
		default:
			e = Subscript{e, o}
		}
	}
	return e
}

func buildAssignable(p any) any {
	var e any = Var{p.(Pair).First.(string)}
	for _, o := range p.(Pair).Second.([]any) {
		switch o.(type) {
		case string:
			e = FieldAccess{e, o.(string)}
		default:
			e = Subscript{e, o}
		}
	}
	return e
}

func buildAssignStatement(p any) any {
	lhs := p.([]any)[0]
	op := p.([]any)[1]
	rhs := p.([]any)[2]
	switch op {
	case "=":
		return Assign{lhs, rhs}
	case "+=":
		return Assign{lhs, BinaryExpr{"+", lhs, rhs}}
	case "-=":
		return Assign{lhs, BinaryExpr{"-", lhs, rhs}}
	case "*=":
		return Assign{lhs, BinaryExpr{"*", lhs, rhs}}
	case "/=":
		return Assign{lhs, BinaryExpr{"/", lhs, rhs}}
	default:
		panic(fmt.Sprintf("invalid assign expression: %v", op))
	}
}

var (
	// 空白字符
	whitespace = Chs(' ', '\t', '\n', '\r')

	// 行注释
	lineComment = Seq(String("//"), Not('\n').Many(), Ch('\n'))

	// 块注释
	blockComment = Seq(String("/*"), Any().ManyUntil(String("*/")), String("*/"))

	// 可忽略元素
	ignorable = OneOf(whitespace, lineComment, blockComment).Many()

	// 字母
	alpha = Range('a', 'z').Or(Range('A', 'Z'))

	// 数字
	digit  = Range('0', '9')
	digits = digit.Many1().Map(join)

	// 下划线
	underline = Ch('_')

	// 整数
	integer = digits.SurroundBy(ignorable)

	// 浮点数
	decimal = Seq(digits, Ch('.'), digits).Map(join).SurroundBy(ignorable)

	// 字符串
	str = Skip(Ch('\'')).And(Not('\'').Many()).Skip(Ch('\'')).Map(join).SurroundBy(ignorable)

	// 布尔值
	boolean = String("true").Or(String("false")).SurroundBy(ignorable)

	// 符号
	assign    = String("=").SurroundBy(ignorable)
	semi      = String(";").SurroundBy(ignorable)
	comma     = String(",").SurroundBy(ignorable)
	colon     = String(":").SurroundBy(ignorable)
	dot       = String(".").SurroundBy(ignorable)
	lp        = String("(").SurroundBy(ignorable)
	rp        = String(")").SurroundBy(ignorable)
	lb        = String("{").SurroundBy(ignorable)
	rb        = String("}").SurroundBy(ignorable)
	ls        = String("[").SurroundBy(ignorable)
	rs        = String("]").SurroundBy(ignorable)
	add       = String("+").SurroundBy(ignorable)
	sub       = String("-").SurroundBy(ignorable)
	mul       = String("*").SurroundBy(ignorable)
	div       = String("/").SurroundBy(ignorable)
	rem       = String("%").SurroundBy(ignorable)
	gt        = String(">").SurroundBy(ignorable)
	get       = String(">=").SurroundBy(ignorable)
	lt        = String("<").SurroundBy(ignorable)
	let       = String("<=").SurroundBy(ignorable)
	equ       = String("==").SurroundBy(ignorable)
	neq       = String("!=").SurroundBy(ignorable)
	and       = String("&&").SurroundBy(ignorable)
	or        = String("||").SurroundBy(ignorable)
	not       = String("!").SurroundBy(ignorable)
	arrow     = String("=>").SurroundBy(ignorable)
	inc       = String("++").SurroundBy(ignorable)
	dec       = String("--").SurroundBy(ignorable)
	addAssign = String("+=").SurroundBy(ignorable)
	subAssign = String("-=").SurroundBy(ignorable)
	mulAssign = String("*=").SurroundBy(ignorable)
	divAssign = String("/=").SurroundBy(ignorable)
	assignOp  = OneOf(assign, addAssign, subAssign, mulAssign, divAssign)

	// 关键字
	import_    = String("import").SurroundBy(ignorable)
	var_       = String("var").SurroundBy(ignorable)
	if_        = String("if").SurroundBy(ignorable)
	else_      = String("else").SurroundBy(ignorable)
	for_       = String("for").SurroundBy(ignorable)
	while_     = String("while").SurroundBy(ignorable)
	break_     = String("break").SurroundBy(ignorable)
	continue_  = String("continue").SurroundBy(ignorable)
	return_    = String("return").SurroundBy(ignorable)
	function_  = String("function").SurroundBy(ignorable)
	undefined_ = String("undefined").SurroundBy(ignorable)
	try_       = String("try").SurroundBy(ignorable)
	catch_     = String("catch").SurroundBy(ignorable)
	finally_   = String("finally").SurroundBy(ignorable)
	throw_     = String("throw").SurroundBy(ignorable)

	// 标识符
	identifier = Seq(OneOf(alpha, underline), OneOf(digit, alpha, underline).Many()).Map(func(r any) any {
		rs := r.([]any)
		return toString(rs[0]) + join(rs[1]).(string)
	}).SurroundBy(ignorable)

	// 整数字面量
	integerLiteral = integer.Map(func(s any) any {
		return Literal{IntegerValue(toInt(s))}
	})

	// 浮点数字面量
	doubleLiteral = decimal.Map(func(s any) any {
		return Literal{DoubleValue(toFloat(s))}
	})

	// 布尔值字面量
	boolLiteral = boolean.Map(func(s any) any {
		return Literal{BoolValue(toBool(s))}
	})

	// 字符串字面量
	stringLiteral = str.Map(func(s any) any {
		return Literal{StringValue(s.(string))}
	})

	// undefined字面量
	undefinedLiteral = undefined_.Map(func(s any) any {
		return Literal{UndefinedValue()}
	})

	// 列表字面量
	listLiteral = Skip(ls).And(exprList).Skip(rs.Fatal()).Map(func(elems any) any {
		return ListLiteral{elems.([]any)}
	})

	// 对象字面量
	fieldPair = OneOf(
		identifier.Skip(colon).And(expr),
		Seq(identifier, lp, idList, rp, lb, stmts, rb).Map(func(r any) any {
			rs := r.([]any)
			return Pair{rs[0], CallableLiteral{rs[2].([]string), Block{rs[5].([]any)}}}
		}),
		identifier.Map(func(id any) any {
			return Pair{id, Var{id.(string)}}
		}),
	)
	fieldList  = SeparateBy(comma, fieldPair).Optional([]any{})
	objLiteral = Skip(lb).And(fieldList).Skip(rb.Fatal()).Map(func(pairs any) any {
		fields := map[string]any{}
		for _, p := range pairs.([]any) {
			fields[p.(Pair).First.(string)] = p.(Pair).Second
		}
		return ObjectLiteral{fields}
	})

	idList = SeparateBy(comma, identifier).Optional([]any{}).Map(func(r any) any {
		ids := []string{}
		for _, id := range r.([]any) {
			ids = append(ids, id.(string))
		}
		return ids
	})

	exprList = SeparateBy(comma, expr).Optional([]any{})

	// 参数列表
	singleParamList = identifier.Map(func(r any) any {
		return []string{r.(string)}
	})
	multiParamList = Skip(lp).And(idList).Skip(rp)
	paramList      = singleParamList.Or(multiParamList)

	stmts = NewParser()

	// 函数字面量
	callableLiteral = paramList.Skip(arrow).And(OneOf(
		// params => {stmts}
		Skip(lb).And(stmts).Skip(rb.Fatal()).Map(func(r any) any {
			return Block{r.([]any)}
		}),
		// params => expr
		expr.Map(func(e any) any {
			return Return{e}
		}),
	)).Map(func(p any) any {
		return CallableLiteral{p.(Pair).First.([]string), p.(Pair).Second}
	})

	// 变量引用
	varRef = identifier.Map(func(s any) any {
		return Var{s.(string)}
	})

	// 下标
	subscript = Skip(ls).And(expr).Skip(rs.Fatal())

	// 字段访问
	fieldAccess = Skip(dot).And(identifier.Fatal())

	// 函数调用
	call = Skip(lp).And(exprList).Skip(rp.Fatal())

	// 表达式
	primaryExpr = NewParser()
	negExpr     = Skip(sub).And(primaryExpr).Map(func(e any) any {
		return UnaryExpr{"-", e}
	})
	notExpr = Skip(not).And(primaryExpr).Map(func(e any) any {
		return UnaryExpr{"!", e}
	})
	bracketExpr        = Skip(lp).And(expr).Skip(rp)
	multiplicativeExpr = primaryExpr.And(OneOf(mul, div, rem).And(primaryExpr).Many()).Map(buildBinaryExpr)
	additiveExpr       = multiplicativeExpr.And(OneOf(add, sub).And(multiplicativeExpr).Many()).Map(buildBinaryExpr)
	relationalExpr     = additiveExpr.And(OneOf(let, lt, get, gt, equ, neq).And(additiveExpr).Many()).Map(buildBinaryExpr)
	andExpr            = relationalExpr.And(and.And(relationalExpr).Many()).Map(buildBinaryExpr)
	expr               = andExpr.And(or.And(andExpr).Many()).Map(buildBinaryExpr)

	// 变量声明
	// var name = expr
	varDeclare = Skip(var_).And(identifier.Fatal()).Skip(assign.Fatal()).And(expr).Map(func(p any) any {
		return VarDeclare{p.(Pair).First.(string), p.(Pair).Second}
	})

	// 函数声明
	// function name(params) {stmts}
	funcDeclare = Seq(function_, identifier.Fatal(), lp.Fatal(), idList, rp.Fatal(), lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
		rs := r.([]any)
		return VarDeclare{rs[1].(string), CallableLiteral{rs[3].([]string), Block{rs[6].([]any)}}}
	})

	// if语句
	ifStmt = Seq(
		// if (expr) {stmts}
		Seq(if_, lp.Fatal(), expr, rp.Fatal(), lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
			rs := r.([]any)
			return Pair{rs[2], Block{rs[5].([]any)}}
		}),
		// (else if (expr) {stmts})*
		Seq(else_, if_, lp.Fatal(), expr, rp.Fatal(), lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
			rs := r.([]any)
			return Pair{rs[3], Block{rs[6].([]any)}}
		}).Many(),
		// else {stmts}
		Seq(else_, lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
			return Block{r.([]any)[2].([]any)}
		}).Optional(Block{[]any{}}),
	).Map(func(r any) any {
		rs := r.([]any)
		cases := []Pair{rs[0].(Pair)}
		for _, p := range rs[1].([]any) {
			cases = append(cases, p.(Pair))
		}
		elseBranch := rs[2].(Block)
		return If{cases, elseBranch}
	})

	// for语句
	// for (init; cond; update) {stmts}
	forStmt = Seq(for_, lp, stmt, semi.Fatal(), expr, semi.Fatal(), stmt, rp.Fatal(), lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
		rs := r.([]any)
		return For{rs[2], rs[4], rs[6], Block{rs[9].([]any)}}
	})

	// while语句
	// while (cond) {stmts}
	whileStmt = Seq(while_, lp.Fatal(), expr, rp.Fatal(), lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
		rs := r.([]any)
		return While{rs[2], Block{rs[5].([]any)}}
	})

	// 代码块
	// {stmts}
	block = Skip(lb).And(stmts).Skip(rb.Fatal()).Map(func(r any) any {
		return Block{r.([]any)}
	})

	// try语句
	tryStmt = Seq(
		// try {stmts}
		Seq(try_, lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
			return Block{r.([]any)[2].([]any)}
		}),
		// catch (id) {stmts}
		Seq(catch_.Fatal(), lp.Fatal(), identifier, rp.Fatal(), lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
			rs := r.([]any)
			return Pair{rs[2], Block{rs[5].([]any)}}
		}),
		// finally {stmts}
		Seq(finally_, lb.Fatal(), stmts, rb.Fatal()).Map(func(r any) any {
			return Block{r.([]any)[2].([]any)}
		}).Optional(Block{[]any{}}),
	).Map(func(r any) any {
		rs := r.([]any)
		return Try{rs[0], rs[1].(Pair).First.(string), rs[1].(Pair).Second, rs[2]}
	})

	// throw语句
	// throw expr
	throwStmt = Skip(throw_).And(expr).Map(func(e any) any {
		return Throw{e}
	})

	// 前置自增
	preInc = Skip(inc).And(assignable).Map(func(r any) any {
		return Assign{r, BinaryExpr{"+", r, Literal{IntegerValue(1)}}}
	})

	// 后置自增
	postInc = assignable.Skip(inc).Map(func(r any) any {
		return Assign{r, BinaryExpr{"+", r, Literal{IntegerValue(1)}}}
	})

	// 前置自减
	preDec = Skip(dec).And(assignable).Map(func(r any) any {
		return Assign{r, BinaryExpr{"-", r, Literal{IntegerValue(1)}}}
	})

	// 后置自减
	postDec = assignable.Skip(dec).Map(func(r any) any {
		return Assign{r, BinaryExpr{"-", r, Literal{IntegerValue(1)}}}
	})

	// break语句
	breakStmt = break_.Map(func(_ any) any {
		return Break{}
	})

	// continue语句
	continueStmt = continue_.Map(func(_ any) any {
		return Continue{}
	})

	// return语句
	returnStmt = Skip(return_).And(expr.Optional(Literal{UndefinedValue()})).Map(func(e any) any {
		return Return{e}
	})

	assignable = identifier.And(OneOf(subscript, fieldAccess).Many()).Map(buildAssignable)

	// 赋值
	assignStmt = Seq(assignable, assignOp, expr).Map(buildAssignStatement)

	// 表达式语句
	exprStmt = expr.Map(func(e any) any {
		return ExprStatement{e}
	})

	// 语句
	stmt = NewParser()

	// 导入声明
	importName = OneOf(digit, alpha, underline, Ch('/')).Many1().SurroundBy(ignorable).Map(join)
	imports    = Skip(import_).And(importName).Many()

	// 程序
	program = imports.And(stmts).Map(func(p any) any {
		strs := []string{}
		for _, s := range p.(Pair).First.([]any) {
			strs = append(strs, s.(string))
		}
		return Program{strs, p.(Pair).Second.([]any)}
	})
)

func init() {
	primaryExpr.Set(OneOf(
		doubleLiteral,
		integerLiteral,
		stringLiteral,
		boolLiteral,
		undefinedLiteral,
		listLiteral,
		objLiteral,
		callableLiteral,
		varRef,
		negExpr,
		notExpr,
		bracketExpr,
	).And(OneOf(call, fieldAccess, subscript).Many()).Map(buildExprElem))

	stmt.Set(OneOf(
		varDeclare,
		funcDeclare,
		ifStmt,
		forStmt,
		whileStmt,
		tryStmt,
		throwStmt,
		block,
		preInc,
		preDec,
		breakStmt,
		continueStmt,
		returnStmt,
		assignStmt,
		postInc,
		postDec,
		exprStmt,
	))

	stmts.Set(stmt.Skip(semi.Optional(nil)).Many())
}

// ParseScript 解析脚本，生成抽象语法树
func ParseScript(script string) Program {
	p, err := program.ParseToEnd(script)
	if err != nil {
		panic(err)
	}
	return p.(Program)
}
