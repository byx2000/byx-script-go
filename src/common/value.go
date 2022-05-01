package common

import (
	"container/list"
	"fmt"
)

const (
	INTEGER = iota
	DOUBLE
	BOOL
	STRING
	LIST
	CALLABLE
	OBJECT
	UNDEFINED
)

type Value struct {
	t    int
	data any
}

type Callable func(args []Value) Value
type Fields map[string]Value

func IntegerValue(v int) Value {
	return Value{INTEGER, v}
}

func DoubleValue(v float64) Value {
	return Value{DOUBLE, v}
}

func BoolValue(v bool) Value {
	return Value{BOOL, v}
}

func StringValue(v string) Value {
	return Value{STRING, v}
}

func ListValue(v *list.List) Value {
	return Value{LIST, v}
}

func CallableValue(v Callable) Value {
	return Value{CALLABLE, v}
}

func ObjectValue(v Fields) Value {
	return Value{OBJECT, v}
}

func UndefinedValue() Value {
	return Value{UNDEFINED, nil}
}

func (v Value) GetType() int {
	return v.t
}

func (v Value) GetData() any {
	return v.data
}

func (v *Value) SetData(data any) {
	v.data = data
}

func (v Value) IsInteger() bool {
	return v.t == INTEGER
}

func (v Value) IsDouble() bool {
	return v.t == DOUBLE
}

func (v Value) IsBool() bool {
	return v.t == BOOL
}

func (v Value) IsString() bool {
	return v.t == STRING
}

func (v Value) IsList() bool {
	return v.t == LIST
}

func (v Value) IsCallable() bool {
	return v.t == CALLABLE
}

func (v Value) IsObject() bool {
	return v.t == OBJECT
}

func (v Value) IsUndefined() bool {
	return v.t == UNDEFINED
}

func (v Value) GetInteger() int {
	return v.data.(int)
}

func (v Value) GetDouble() float64 {
	return v.data.(float64)
}

func (v Value) GetBool() bool {
	return v.data.(bool)
}

func (v Value) GetString() string {
	return v.data.(string)
}

func (v Value) GetList() *list.List {
	return v.data.(*list.List)
}

func (v Value) GetCallable() Callable {
	return v.data.(Callable)
}

func (v Value) GetObject() Fields {
	return v.data.(Fields)
}

func (v Value) String() string {
	if v.t == UNDEFINED {
		return "undefined"
	}
	return fmt.Sprintf("%v", v.data)
}
