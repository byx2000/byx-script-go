package parser

import . "byx-script-go/src/common"

// 表达式

// Literal 字面量（整数、浮点数、布尔值）
type Literal struct {
	V Value
}

// ListLiteral 列表字面量
type ListLiteral struct {
	Elems []any
}

// CallableLiteral 函数字面量
type CallableLiteral struct {
	Params []string
	Body   any
}

// ObjectLiteral 对象字面量
type ObjectLiteral struct {
	Fields map[string]any
}

// Var 变量引用
type Var struct {
	VarName string
}

// UnaryExpr 一元表达式
type UnaryExpr struct {
	Op   string
	Expr any
}

// BinaryExpr 二元表达式
type BinaryExpr struct {
	Op  string
	Lhs any
	Rhs any
}

// Call 函数调用
type Call struct {
	Callee any
	Args   []any
}

// FieldAccess 字段访问
type FieldAccess struct {
	Expr any
	Name string
}

// Subscript 下标访问
type Subscript struct {
	Expr any
	Sub  any
}

// 语句

// VarDeclare 变量定义
type VarDeclare struct {
	VarName string
	Value   any
}

// Assign 赋值
type Assign struct {
	Lhs any
	Rhs any
}

// Block 语句块
type Block struct {
	Stmts []any
}

// If if语句
type If struct {
	Cases      []Pair
	ElseBranch any
}

// While while循环
type While struct {
	Cond any
	Body any
}

// For for循环
type For struct {
	Init   any
	Cond   any
	Update any
	Body   any
}

// Continue continue语句
type Continue struct {
}

// Break break语句
type Break struct {
}

// Return return语句
type Return struct {
	RetVal any
}

// ExprStatement 表达式语句
type ExprStatement struct {
	Expr any
}

// Throw throw语句
type Throw struct {
	Expr any
}

// Try try语句
type Try struct {
	TryBranch     any
	CatchVar      string
	CatchBranch   any
	FinallyBranch any
}

type Program struct {
	Imports []string
	Stmts   []any
}
