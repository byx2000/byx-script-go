package common

import (
	"container/list"
	"fmt"
)

// Value类型常量
const (
	typeInteger = iota
	typeDouble
	typeBool
	typeString
	typeList
	typeCallable
	typeObject
	typeUndefined
)

// Value ByxScript值的底层表示
type Value struct {
	t    int // 类型
	data any // 数据
}

type Callable func(args []Value) Value
type Fields map[string]Value

var undefinedValue = Value{typeUndefined, nil}
var trueValue = Value{typeBool, true}
var falseValue = Value{typeBool, false}

// IntegerValue 创建整数值
func IntegerValue(v int) Value {
	return Value{typeInteger, v}
}

// DoubleValue 创建浮点数值
func DoubleValue(v float64) Value {
	return Value{typeDouble, v}
}

// BoolValue 创建布尔值
func BoolValue(v bool) Value {
	if v {
		return trueValue
	} else {
		return falseValue
	}
}

// StringValue 创建字符串值
func StringValue(v string) Value {
	return Value{typeString, v}
}

// ListValue 创建列表值
func ListValue(v *list.List) Value {
	return Value{typeList, v}
}

// CallableValue 创建可调用值
func CallableValue(v Callable) Value {
	return Value{typeCallable, v}
}

// ObjectValue 创建对象值
func ObjectValue(v Fields) Value {
	return Value{typeObject, v}
}

// UndefinedValue 创建undefined值
func UndefinedValue() Value {
	return undefinedValue
}

// GetType 获取值类型
func (v Value) GetType() int {
	return v.t
}

// GetData 获取值底层数据
func (v Value) GetData() any {
	return v.data
}

// IsInteger 当前值是否为整数
func (v Value) IsInteger() bool {
	return v.t == typeInteger
}

// IsDouble 当前值是否为浮点数
func (v Value) IsDouble() bool {
	return v.t == typeDouble
}

// IsBool 当前值是否为布尔
func (v Value) IsBool() bool {
	return v.t == typeBool
}

// IsString 当前值是否为字符串
func (v Value) IsString() bool {
	return v.t == typeString
}

// IsList 当前值是否为列表
func (v Value) IsList() bool {
	return v.t == typeList
}

// IsCallable 当前值是否可调用
func (v Value) IsCallable() bool {
	return v.t == typeCallable
}

// IsObject 当前值是否为对象
func (v Value) IsObject() bool {
	return v.t == typeObject
}

// IsUndefined 当前值是否为undefined
func (v Value) IsUndefined() bool {
	return v.t == typeUndefined
}

// GetInteger 获取整数值
func (v Value) GetInteger() int {
	return v.data.(int)
}

// GetDouble 获取浮点数值
func (v Value) GetDouble() float64 {
	return v.data.(float64)
}

// GetBool 获取布尔值
func (v Value) GetBool() bool {
	return v.data.(bool)
}

// GetString 获取字符串值
func (v Value) GetString() string {
	return v.data.(string)
}

// GetList 获取列表值
func (v Value) GetList() *list.List {
	return v.data.(*list.List)
}

// GetCallable 获取可调用值
func (v Value) GetCallable() Callable {
	return v.data.(Callable)
}

// GetFields 获取对象属性
func (v Value) GetFields() Fields {
	return v.data.(Fields)
}

func (v Value) String() string {
	if v.t == typeUndefined {
		return "undefined"
	}
	return fmt.Sprintf("%v", v.data)
}
