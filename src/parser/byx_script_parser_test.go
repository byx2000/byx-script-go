package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerLiteral(t *testing.T) {
	r, _ := integerLiteral.ParseToEnd("  123  ")
	assert.Equal(t, Literal{123}, r)
}

func TestDoubleLiteral(t *testing.T) {
	r, _ := doubleLiteral.ParseToEnd("  23.75  ")
	assert.Equal(t, Literal{23.75}, r)
}

func TestBoolLiteral(t *testing.T) {
	r, _ := boolLiteral.ParseToEnd(" true  ")
	assert.Equal(t, Literal{true}, r)
	r, _ = boolLiteral.ParseToEnd(" false ")
	assert.Equal(t, Literal{false}, r)
}

func TestStringLiteral(t *testing.T) {
	r, _ := stringLiteral.ParseToEnd(" 'hello'  ")
	assert.Equal(t, Literal{"hello"}, r)
	r, _ = stringLiteral.ParseToEnd(" ''  ")
	assert.Equal(t, Literal{""}, r)
}

func TestUndefinedLiteral(t *testing.T) {
	r, _ := undefinedLiteral.ParseToEnd(" undefined  ")
	assert.Equal(t, Literal{nil}, r)
}

func TestListLiteral(t *testing.T) {
	r, _ := listLiteral.ParseToEnd(" [1, 2, 3] ")
	assert.Equal(t, ListLiteral{[]any{Literal{1}, Literal{2}, Literal{3}}}, r)
	r, _ = listLiteral.ParseToEnd(" [123] ")
	assert.Equal(t, ListLiteral{[]any{Literal{123}}}, r)
	r, _ = listLiteral.ParseToEnd(" [] ")
	assert.Equal(t, ListLiteral{[]any{}}, r)
}

func TestCallableLiteral(t *testing.T) {
	r, _ := callableLiteral.ParseToEnd("() => 100")
	assert.Equal(t, CallableLiteral{[]string{}, Return{Literal{100}}}, r)
	r, _ = callableLiteral.ParseToEnd("a => a + 1")
	assert.Equal(t, CallableLiteral{[]string{"a"}, Return{BinaryExpr{"+", Var{"a"}, Literal{1}}}}, r)
	r, _ = callableLiteral.ParseToEnd("(a, b) => a + b")
	assert.Equal(t, CallableLiteral{[]string{"a", "b"}, Return{BinaryExpr{"+", Var{"a"}, Var{"b"}}}}, r)
	r, _ = callableLiteral.ParseToEnd("(a, b) => {var c = a * b return c}")
	assert.Equal(t, CallableLiteral{[]string{"a", "b"}, Block{[]any{VarDeclare{"c", BinaryExpr{"*", Var{"a"}, Var{"b"}}}, Return{Var{"c"}}}}}, r)
}

func TestParamList(t *testing.T) {
	r, _ := paramList.ParseToEnd(" ( a , b , nums , c ) ")
	assert.Equal(t, []string{"a", "b", "nums", "c"}, r)
	r, _ = paramList.ParseToEnd(" x ")
	assert.Equal(t, []string{"x"}, r)
}

func TestSubscript(t *testing.T) {
	r, _ := subscript.ParseToEnd(" [ 10 ] ")
	assert.Equal(t, Literal{10}, r)
	r, _ = subscript.ParseToEnd(" [ 2 + 3 ] ")
	assert.Equal(t, BinaryExpr{"+", Literal{2}, Literal{3}}, r)
}

func TestFieldAccess(t *testing.T) {
	r, _ := fieldAccess.ParseToEnd(" .name ")
	assert.Equal(t, "name", r)
}

func TestCall(t *testing.T) {
	r, _ := call.ParseToEnd(" (1, 2, 3) ")
	assert.Equal(t, []any{Literal{1}, Literal{2}, Literal{3}}, r)
}

func TestVarRef(t *testing.T) {
	r, _ := varRef.ParseToEnd(" nums ")
	assert.Equal(t, Var{"nums"}, r)
}

func TestPrimaryExpr(t *testing.T) {
	r, _ := primaryExpr.ParseToEnd(" 123 ")
	assert.Equal(t, Literal{123}, r)
	r, _ = primaryExpr.ParseToEnd(" 34.56 ")
	assert.Equal(t, Literal{34.56}, r)
	r, _ = primaryExpr.ParseToEnd(" 'hello' ")
	assert.Equal(t, Literal{"hello"}, r)
	r, _ = primaryExpr.ParseToEnd(" num ")
	assert.Equal(t, Var{"num"}, r)
	r, _ = primaryExpr.ParseToEnd(" [1, 2, 3] ")
	assert.Equal(t, ListLiteral{[]any{Literal{1}, Literal{2}, Literal{3}}}, r)
	r, _ = primaryExpr.ParseToEnd(" -123 ")
	assert.Equal(t, UnaryExpr{"-", Literal{123}}, r)
	r, _ = primaryExpr.ParseToEnd(" !false ")
	assert.Equal(t, UnaryExpr{"!", Literal{false}}, r)
	r, _ = primaryExpr.ParseToEnd(" (2 + 3) ")
	assert.Equal(t, BinaryExpr{"+", Literal{2}, Literal{3}}, r)
	r, _ = primaryExpr.ParseToEnd(" a.b ")
	assert.Equal(t, FieldAccess{Var{"a"}, "b"}, r)
	r, _ = primaryExpr.ParseToEnd(" a[100] ")
	assert.Equal(t, Subscript{Var{"a"}, Literal{100}}, r)
	r, _ = primaryExpr.ParseToEnd(" fun(123, 3.14) ")
	assert.Equal(t, Call{Var{"fun"}, []any{Literal{123}, Literal{3.14}}}, r)
	r, _ = primaryExpr.ParseToEnd(" a.b[c].d(e, f) ")
	assert.Equal(t, Call{FieldAccess{Subscript{FieldAccess{Var{"a"}, "b"}, Var{"c"}}, "d"}, []any{Var{"e"}, Var{"f"}}}, r)
}

func TestExpr(t *testing.T) {
	r, _ := expr.ParseToEnd(" 123 ")
	assert.Equal(t, 123, r.(Literal).Value)
	r, _ = expr.ParseToEnd(" 3.14 ")
	assert.Equal(t, 3.14, r.(Literal).Value)
	r, _ = expr.ParseToEnd(" true ")
	assert.Equal(t, true, r.(Literal).Value)
	r, _ = expr.ParseToEnd(" 'hello' ")
	assert.Equal(t, "hello", r.(Literal).Value)
	r, _ = expr.ParseToEnd(" 123 + 45.6 * 789")
	assert.Equal(t, BinaryExpr{"+", Literal{123}, BinaryExpr{"*", Literal{45.6}, Literal{789}}}, r)
}

func TestVarDeclare(t *testing.T) {
	r, _ := varDeclare.ParseToEnd(" var num = 1 + 2 ")
	assert.Equal(t, VarDeclare{"num", BinaryExpr{"+", Literal{1}, Literal{2}}}, r)
}

func TestAssign(t *testing.T) {
	r, _ := assignStmt.ParseToEnd(" num = 123 ")
	assert.Equal(t, Assign{Var{"num"}, Literal{123}}, r)
	r, _ = assignStmt.ParseToEnd(" a.b = 3.14 + 2 ")
	assert.Equal(t, Assign{FieldAccess{Var{"a"}, "b"}, BinaryExpr{"+", Literal{3.14}, Literal{2}}}, r)
	r, _ = assignStmt.ParseToEnd(" a[0].b = 100 ")
	assert.Equal(t, Assign{FieldAccess{Subscript{Var{"a"}, Literal{0}}, "b"}, Literal{100}}, r)
	r, _ = assignStmt.ParseToEnd(" arr[100] = 'hello' ")
	assert.Equal(t, Assign{Subscript{Var{"arr"}, Literal{100}}, Literal{"hello"}}, r)
	r, _ = assignStmt.ParseToEnd(" arr.b[100] = 3.14 ")
	assert.Equal(t, Assign{Subscript{FieldAccess{Var{"arr"}, "b"}, Literal{100}}, Literal{3.14}}, r)
}

func TestIf(t *testing.T) {
	r, _ := ifStmt.ParseToEnd("if (true) {a = 123} else {a = 456}")
	fmt.Println(r)
}

func TestReturn(t *testing.T) {
	r, _ := returnStmt.ParseToEnd(" return 123 ")
	assert.Equal(t, Return{Literal{123}}, r)
	r, _ = returnStmt.ParseToEnd(" return nums ")
	assert.Equal(t, Return{Var{"nums"}}, r)
	r, _ = returnStmt.ParseToEnd(" return ")
	assert.Equal(t, Return{Literal{nil}}, r)
}
