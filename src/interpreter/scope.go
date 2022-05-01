package interpreter

import (
	. "byx-script-go/src/common"
	"fmt"
)

// Scope 作用域
type Scope struct {
	vars map[string]Value
	next *Scope
}

// NewEmptyScope 创建空的作用域
func NewEmptyScope() Scope {
	return Scope{
		vars: map[string]Value{},
		next: nil,
	}
}

// NewScope 创建作用域
func NewScope(next Scope) Scope {
	return Scope{
		vars: map[string]Value{},
		next: &next,
	}
}

// DeclareVar 定义变量
func (s Scope) DeclareVar(varName string, value Value) {
	if _, exist := s.vars[varName]; exist {
		panic(fmt.Sprintf("var already exist: %s", varName))
	}
	s.vars[varName] = value
}

// SetVar 设置变量的值
func (s Scope) SetVar(varName string, value Value) {
	cur := &s
	for {
		if cur == nil {
			break
		}
		if _, exist := cur.vars[varName]; exist {
			cur.vars[varName] = value
			return
		}
		cur = cur.next
	}
	panic(fmt.Sprintf("var not exist: %s", varName))
}

// GetVar 获取变量的值
func (s Scope) GetVar(varName string) Value {
	cur := &s
	for {
		if cur == nil {
			break
		}
		if v, exist := cur.vars[varName]; exist {
			return v
		}
		cur = cur.next
	}
	panic(fmt.Sprintf("var not exist: %s", varName))
}
