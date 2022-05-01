package interpreter

import (
	. "byx-script-go/src/common"
	"container/list"
	"fmt"
	"hash/adler32"
	"math"
	"strings"
	"unicode/utf8"
)

func printValue(v Value, deep bool) {
	if v.IsInteger() || v.IsDouble() || v.IsBool() || v.IsString() {
		fmt.Print(v.GetData())
	} else if v.IsList() {
		if deep {
			fmt.Print("[")
			lst := v.GetList()
			for e := lst.Front(); e != nil; e = e.Next() {
				printValue(e.Value.(Value), false)
				if e.Next() != nil {
					fmt.Print(", ")
				}
			}
			fmt.Println("]")
		} else {
			fmt.Print("[...]")
		}
	} else if v.IsObject() {
		if deep {
			fmt.Print("{")
			i := 0
			fields := v.GetObject()
			for k, v := range fields {
				fmt.Print(k + ": ")
				printValue(v, false)
				if i != len(fields)-1 {
					fmt.Print(", ")
				}
				i++
			}
			fmt.Print("}")
		} else {
			fmt.Print("{...}")
		}
	} else if v.IsUndefined() {
		fmt.Print("undefined")
	} else {
		panic(fmt.Sprintf("unknown print value: %s", v.String()))
	}
}

func printValues(values []Value) {
	for i, v := range values {
		printValue(v, true)
		if i != len(values)-1 {
			fmt.Print(" ")
		}
	}
}

var Print = CallableValue(func(args []Value) Value {
	printValues(args)
	return UndefinedValue()
})

var Println = CallableValue(func(args []Value) Value {
	printValues(args)
	fmt.Println()
	return UndefinedValue()
})

func StringLength(s string) Value {
	return CallableValue(func(args []Value) Value {
		return IntegerValue(utf8.RuneCountInString(s))
	})
}

func StringSubstring(s string) Value {
	return CallableValue(func(args []Value) Value {
		start := args[0].GetInteger()
		end := args[1].GetInteger()
		return StringValue(string([]rune(s)[start:end]))
	})
}

func StringConcat(s string) Value {
	return CallableValue(func(args []Value) Value {
		return StringValue(s + args[0].GetString())
	})
}

func StringCharAt(s string) Value {
	return CallableValue(func(args []Value) Value {
		index := args[0].GetInteger()
		return StringValue(string([]rune(s)[index : index+1]))
	})
}

func StringCodeAt(s string) Value {
	return CallableValue(func(args []Value) Value {
		index := args[0].GetInteger()
		return IntegerValue(int([]rune(s)[index]))
	})
}

func StringToUpper(s string) Value {
	return CallableValue(func(args []Value) Value {
		return StringValue(strings.ToUpper(s))
	})
}

func StringToLower(s string) Value {
	return CallableValue(func(args []Value) Value {
		return StringValue(strings.ToLower(s))
	})
}

func ListLength(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		return IntegerValue(lst.Len())
	})
}

func ListAddFirst(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		lst.PushFront(args[0])
		return UndefinedValue()
	})
}

func ListRemoveFirst(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		return lst.Remove(lst.Front()).(Value)
	})
}

func ListAddLast(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		lst.PushBack(args[0])
		return UndefinedValue()
	})
}

func ListRemoveLast(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		return lst.Remove(lst.Back()).(Value)
	})
}

func ListInsert(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		index := args[0].GetInteger()
		v := args[1]
		if index == lst.Len() {
			lst.PushBack(v)
			return UndefinedValue()
		}
		for e := lst.Front(); e != nil; e = e.Next() {
			if index == 0 {
				lst.InsertBefore(v, e)
				return UndefinedValue()
			}
			index--
		}
		return UndefinedValue()
	})
}

func ListRemove(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		index := args[0].GetInteger()
		for e := lst.Front(); e != nil; e = e.Next() {
			if index == 0 {
				return lst.Remove(e).(Value)
			}
			index--
		}
		return UndefinedValue()
	})
}

func ListCopy(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		newList := list.New()
		for e := lst.Front(); e != nil; e = e.Next() {
			newList.PushBack(e.Value)
		}
		return ListValue(newList)
	})
}

func ListIsEmpty(lst *list.List) Value {
	return CallableValue(func(args []Value) Value {
		return BoolValue(lst.Len() == 0)
	})
}

var IsInteger = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsInteger())
})

var IsDouble = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsDouble())
})

var IsBool = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsBool())
})

var IsString = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsString())
})

var IsList = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsList())
})

var IsCallable = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsCallable())
})

var IsObject = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsObject())
})

var IsUndefined = CallableValue(func(args []Value) Value {
	return BoolValue(args[0].IsUndefined())
})

var Hashcode = CallableValue(func(args []Value) Value {
	bytes := []byte(fmt.Sprintf("%v", args[0].GetData()))
	return IntegerValue(int(adler32.Checksum(bytes)))
})

var Pow = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Pow(args[0].GetDouble(), args[1].GetDouble()))
})

var Sin = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Sin(args[0].GetDouble()))
})

var Cos = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Cos(args[0].GetDouble()))
})

var Tan = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Tan(args[0].GetDouble()))
})

var Exp = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Exp(args[0].GetDouble()))
})

var Ln = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Log(args[0].GetDouble()))
})

var Log10 = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Log10(args[0].GetDouble()))
})

var Sqrt = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Sqrt(args[0].GetDouble()))
})

var Round = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Round(args[0].GetDouble()))
})

var Ceil = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Ceil(args[0].GetDouble()))
})

var Floor = CallableValue(func(args []Value) Value {
	return DoubleValue(math.Floor(args[0].GetDouble()))
})
