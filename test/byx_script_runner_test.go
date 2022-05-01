package test

import (
	. "byx-script-go/src/interpreter"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	verify(t, `
	println(123 + 456)
	println(123 + 3.14)
	println(12.34 + 555)
	println(12.34 + 56.78)
	println('hello ' + 'world!')
	println('hello ' + 123)
	println(123 + ' hello')
	println('world ' + 3.14)
	println(3.14 + ' world')
	println('abc ' + true)
	println(false + ' abc')
	println(undefined + ' xyz')
	println('xyz ' + undefined)`, getOutput(func() {
		fmt.Println(123 + 456)
		fmt.Println(123 + 3.14)
		fmt.Println(12.34 + 555)
		fmt.Println(12.34 + 56.78)
		fmt.Println("hello " + "world!")
		fmt.Println("hello " + strconv.Itoa(123))
		fmt.Println(strconv.Itoa(123) + " hello")
		fmt.Println(fmt.Sprintf("world %v", 3.14))
		fmt.Println(fmt.Sprintf("%v world", 3.14))
		fmt.Println("abc true")
		fmt.Println("false abc")
		fmt.Println("undefined xyz")
		fmt.Println("xyz undefined")
	}))
}

func TestSub(t *testing.T) {
	verify(t, `
	println(532 - 34)
	println(3.14 - 12)
	println(12 - 7.78)
	println(56.78 - 12.34)`, getOutput(func() {
		fmt.Println(532 - 34)
		fmt.Println(3.14 - 12)
		fmt.Println(12 - 7.78)
		fmt.Println(56.78 - 12.34)
	}))
}

func TestMul(t *testing.T) {
	verify(t, `
	println(12 * 34)
	println(12 * 3.4)
	println(0.12 * 34)
	println(12.34 * 56.78)`, getOutput(func() {
		fmt.Println(12 * 34)
		fmt.Println(12 * 3.4)
		fmt.Println(0.12 * 34)
		fmt.Println(12.34 * 56.78)
	}))
}

func TestDiv(t *testing.T) {
	verify(t, `
	println(5 / 2)
	println(12 / 3.4)
	println(1.2 / 10)
	println(56.78 / 12.34)`, getOutput(func() {
		fmt.Println(5 / 2)
		fmt.Println(12 / 3.4)
		fmt.Println(1.2 / 10)
		fmt.Println(56.78 / 12.34)
	}))
}

func TestRem(t *testing.T) {
	verify(t, `
	println(12 % 3)
	println(12 % 5)
	println(3 % 7)
	println(6 % 3)`, getOutput(func() {
		fmt.Println(12 % 3)
		fmt.Println(12 % 5)
		fmt.Println(3 % 7)
		fmt.Println(6 % 3)
	}))
}

func TestGreaterThan(t *testing.T) {
	verify(t, `
	println(100 > 50)
	println(100 > 100)
	println(3.14 > 50)
	println(3.14 > 1)
	println(3.14 > 456.23)
	println(3.14 > 3.14)
	println('banana' > 'apple')
	println('apple' > 'banana')
	println('apple' > 'apple')`, `
	true
	false
	false
	true
	false
	false
	true
	false
	false`)
}

func TestGreaterEqualThan(t *testing.T) {
	verify(t, `
	println(100 >= 50)
	println(100 >= 100)
	println(3.14 >= 50)
	println(3.14 >= 1)
	println(3.14 >= 456.23)
	println(3.14 >= 3.14)
	println('banana' >= 'apple')
	println('apple' >= 'banana')
	println('apple' >= 'apple')`, `
	true
	true
	false
	true
	false
	true
	true
	false
	true`)
}

func TestLessThan(t *testing.T) {
	verify(t, `
	println(100 < 50)
	println(100 < 100)
	println(3.14 < 50)
	println(3.14 < 1)
	println(3.14 < 456.23)
	println(3.14 < 3.14)
	println('banana' < 'apple')
	println('apple' < 'banana')
	println('apple' < 'apple')`, `
	false
	false
	true
	false
	true
	false
	false
	true
	false`)
}

func TestLessEqualThan(t *testing.T) {
	verify(t, `
	println(100 <= 50)
	println(100 <= 100)
	println(3.14 <= 50)
	println(3.14 <= 1)
	println(3.14 <= 456.23)
	println(3.14 <= 3.14)
	println('banana' <= 'apple')
	println('apple' <= 'banana')
	println('apple' <= 'apple')`, `
	false
	true
	true
	false
	true
	true
	false
	true
	true`)
}

func TestEqual(t *testing.T) {
	verify(t, `
	println(123 == 123)
	println(12.34 == 12.34)
	println(true == true)
	println(false == false)
	println('apple' == 'apple')
	println(123 == 45)
	println(3.14 == 12.56)
	println(true == false)
	println('apple' == 'banana')
	println({a: 123, b: 'hello'} == {a: 123, b: 'hello'})
	
	var a = {a: 123, b: 'hello'}
	var b = a
	var c = {a: 123, b: 'hello'}
	println(a == b)
	println(a == c)`, `
	true
	true
	true
	true
	true
	false
	false
	false
	false
	false
	true
	false`)
}

func TestNotEqual(t *testing.T) {
	verify(t, `
	println(123 != 123)
	println(12.34 != 12.34)
	println(true != true)
	println(false != false)
	println('apple' != 'apple')
	println(123 != 45)
	println(3.14 != 12.56)
	println(true != false)
	println('apple' != 'banana')
	println({a: 123, b: 'hello'} != {a: 123, b: 'hello'})
	
	var a = {a: 123, b: 'hello'}
	var b = a
	var c = {a: 123, b: 'hello'}
	println(a != b)
	println(a != c)`, `
	false
	false
	false
	false
	false
	true
	true
	true
	true
	true
	false
	true`)
}

func TestAnd(t *testing.T) {
	verify(t, `
	println(true && true)
	println(true && false)
	println(false && true)
	println(false && false)`, `
	true
	false
	false
	false`)

	verify(t, `
	var x = 0
	var y = 0
	function f1() {
		x = x + 1
		return true
	}
	function f2() {
		y = y + 1
		return false
	}
					
	x = 0
	y = 0
	var b = f2() && f1()
	println(b)
	println(x)
	println(y)`, `
	false
	0
	1`)
}

func TestOr(t *testing.T) {
	verify(t, `
	println(true || true)
	println(true || false)
	println(false || true)
	println(false || false)`, `
	true
	true
	true
	false`)

	verify(t, `
	var x = 0
	var y = 0
	function f1() {
		x = x + 1
		return true
	}
	function f2() {
		y = y + 1
		return false
	}

	x = 0
	y = 0
	var b = f1() || f2()
	println(b)
	println(x)
	println(y)
	`, `
	true
	1
	0`)
}

func TestNot(t *testing.T) {
	verify(t, `
	println(!true)
	println(!false)`, `
	false
	true`)
}

func TestExpr(t *testing.T) {
	verify(t, `
	println(2 + 3*5)
	println(12 + 34 + 56 * 78 *90)
	println((2+3) * 4 / (9-7))
	println(-100)
	println(-5 + 7)
	println(-(5 + 7))
	println(-3.14)
	println(-12.34-67.5)
	println(-(12.34-67.5))
	
	println(!false || true)
	println(!(false || true))
	println(!true && false)
	println(!(true && false))
	println(true && !false)
	println(false || !true)
	println(false && false || true)
	println(false && (false || true))
	println(true || true && false)
	println((true || true) && false)
	println(true && true && true)
	println(true && true && false)
	println(true || false || true)
	println(false || false || false)`, getOutput(func() {
		fmt.Println(2 + 3*5)
		fmt.Println(12 + 34 + 56*78*90)
		fmt.Println((2 + 3) * 4 / (9 - 7))
		fmt.Println(-100)
		fmt.Println(-5 + 7)
		fmt.Println(-(5 + 7))
		fmt.Println(-3.14)
		fmt.Println(-12.34 - 67.5)
		fmt.Println(-(12.34 - 67.5))

		fmt.Println(true)
		fmt.Println(false)
		fmt.Println(false)
		fmt.Println(true)
		fmt.Println(true)
		fmt.Println(false)
		fmt.Println(true)
		fmt.Println(false)
		fmt.Println(true)
		fmt.Println(false)
		fmt.Println(true)
		fmt.Println(false)
		fmt.Println(true)
		fmt.Println(false)
	}))
}

func TestString(t *testing.T) {
	// length
	verify(t, `
	println(''.length())
	println('abc'.length())
	println('hello，你好'.length())`, `
	0
	3
	8`)

	// substring
	verify(t, `
	var s = 'hello';
	println(s.substring(1, 4))
	println('你好世界'.substring(1, 3))`, `
	ell
	好世`)

	// concat
	verify(t, `
	println('abc'.concat('defg'))`, `
	abcdefg`)

	// charAt
	verify(t, `
	var s = 'abc'
	println(s.charAt(0), s.charAt(1), s.charAt(2))
	println(s.charAt(1) == 'b')
	println('你好'.charAt(0))`, `
	a b c
	true
	你`)

	// codeAt
	verify(t, `
	var s = 'abc'
	println(s.codeAt(0), s.codeAt(1), s.codeAt(2))
	println('你好'.codeAt(1) == '好'.codeAt(0))`, `
	97 98 99
	true`)

	// toUpper
	verify(t, `
	println('abc'.toUpper())
	println('Abc'.toUpper())
	println('ABC'.toUpper())`, `
	ABC
	ABC
	ABC`)

	// toLower
	verify(t, `
	println('abc'.toLower())
	println('Abc'.toLower())
	println('ABC'.toLower())`, `
	abc
	abc
	abc`)

	// index
	verify(t, `
	var s = 'abc'
	println(s[0], s[1], s[2])
	println(s[1] == 'b')
	println('hello，世界'[6])`, `
	a b c
	true
	世`)
}

func TestList(t *testing.T) {
	// length
	verify(t, `
	var arr1 = []
	println(arr1.length())
	var arr2 = [1, 2, 3, 4]
	println(arr2.length())
	println([1, 2, 3].length())`, `
	0
	4
	3`)

	// addFirst
	verify(t, `
	var arr = [1, 2, 3]
	println(arr.length())
	arr.addFirst(4)
	arr.addFirst(5)
	println(arr.length())
	arr.addFirst(3.14)
	arr.addFirst('hello')
	println(arr.length())
	println(arr)`, `
	3
	5
	7
	[hello, 3.14, 5, 4, 1, 2, 3]`)

	// removeFirst
	verify(t, `
	var nums = [1, 2, 3, 4, 5]
	println(nums.removeFirst())
	println(nums)`, `
	1
	[2, 3, 4, 5]`)

	// addLast
	verify(t, `
	var arr = [1, 2, 3]
	println(arr.length())
	arr.addLast(4)
	arr.addLast(5)
	println(arr.length())
	arr.addLast(3.14)
	arr.addLast('hello')
	println(arr.length())
	println(arr)`, `
	3
	5
	7
	[1, 2, 3, 4, 5, 3.14, hello]`)

	// removeLast
	verify(t, `
	var nums = [1, 2, 3, 4, 5]
	println(nums.removeLast())
	println(nums)`, `
	5
	[1, 2, 3, 4]`)

	// insert
	verify(t, `
	var list = [1, 2, 3, 4, 5]
	list.insert(0, 100)
	list.insert(3, 'hello')
	list.insert(7, 3.14)
	println(list)`, `
	[100, 1, 2, hello, 3, 4, 5, 3.14]`)

	// remove
	verify(t, `
	var list = [1, 2, 3, 4, 5]
	println(list.remove(2))
	println(list)`, `
	3
	[1, 2, 4, 5]`)

	// copy
	verify(t, `
	var list1 = [1, 2, 3, 4, 5]
	var list2 = list1.copy()
	list1[2] = 100
	list2[3] = 200
	println(list1)
	println(list2)`, `
	[1, 2, 100, 4, 5]
	[1, 2, 3, 200, 5]`)

	// isEmpty
	verify(t, `
	var a1 = [1, 2, 3]	
	println(a1.isEmpty())
	var a2 = []
	println(a2.isEmpty())`, `
	false
	true`)

	verify(t, `
	var nums = []
	for (var i = 1; i <= 100; i = i + 1) {
		nums.addLast(i * i)
	}
	var s = 0
	for (var i = 0; i < nums.length(); i = i + 1) {
		s = s + nums[i]
	}
	println(s)`, `
	338350`)

	verify(t, `
	var list = []
	for (var i = 10; i <= 30; i += 10) {
		list.addLast([])
		for (var j = 0; j < 4; j = j + 1) {
			list[i / 10 - 1].addLast(i + j)
		}
	}
	
	for (var i = 0; i < list.length(); ++i) {
		for (var j = 0; j < list[i].length(); ++j) {
			print(list[i][j] + ' ')
		}
		println()
	}`, `
	10 11 12 13
	20 21 22 23
	30 31 32 33`)
}

func TestObject(t *testing.T) {
	verify(t, `
	var age = 21
	var obj = {
		name: 'Tom',
		age,
		add: (a, b) => a + b,
		sub(a, b) {
			return a - b		
		}
	}
	println(obj.name)
	println(obj.age)
	println(obj.add(2, 3))
	println(obj.sub(200, 100))
	`, `
	Tom
	21
	5
	100`)
}

func TestScope(t *testing.T) {
	verify(t, `
	var i = 100 + 2*3;
	i = 123;
	var j = i + 1;
	{
		var k = 3
		i = k
	}
	println(i)
	println(j)`, `
	3
	124`)
}

func TestIf(t *testing.T) {
	verify(t, `
	if (true) {}
	if (true) {} else {}
	if (true) {} else if (false) {} else if (true) {} else {}`, ``)

	verify(t, `
	var i = 123
	if (i > 200) {
		i = 456
	}
	println(i)
	`, `
	123`)

	verify(t, `
	var i = 123
	if (200 > i) {
		i = 456
	}
	println(i)`, `
	456`)

	verify(t, `
	var i = 100
	if (i < 25 || !(i < 50)) {
		i = 200
	}
	println(i)`, `
	200`)

	verify(t, `
	var i = 123
	var j = 456
	if (i < 200 && j > 300) {
		i = 1001
		j = 1002
	} else {
		i = 1003
		j = 1004
	}
	println(i, j)`, `
	1001 1002`)

	verify(t, `
	var i = 123
	var j = 456
	if (i > 200 && j > 300) {
		i = 1001
		j = 1002
	} else {
		i = 1003
		j = 1004
	}
	println(i, j)`, `
	1003 1004`)

	verify(t, `
	function getLevel(score) {
		if (85 < score && score <= 100) {
			return 'excellent'
		} else if (75 < score && score <= 85) {
			return 'good'
		} else if (60 < score && score <= 75) {
			return 'pass'
		} else {
			return 'failed'
		}
	}
	
	println(getLevel(92))
	println(getLevel(73))
	println(getLevel(81))
	println(getLevel(50))`, `
	excellent
	pass
	good
	failed`)
}

func TestInc(t *testing.T) {
	verify(t, `
	var i = 100
	var obj = {a: 1, b: {x: 20}, c: [1, 2, 3]}
	
	i++
	++i
	obj.a++
	++obj.a
	obj.c[0]++
	++obj.c[1]
	
	println(i)
	println(obj.a)
	println(obj.c[0])
	println(obj.c[1])`, `
	102
	3
	2
	3`)
}

func TestDec(t *testing.T) {
	verify(t, `
	var i = 100
	var obj = {a: 1, b: {x: 20}, c: [1, 2, 3]}
	
	i--
	--i
	obj.a--
	--obj.a
	obj.c[0]--
	--obj.c[1]
	
	println(i)
	println(obj.a)
	println(obj.c[0])
	println(obj.c[1])`, `
	98
	-1
	0
	1`)
}

func TestFor(t *testing.T) {
	verify(t, `
	for (var i = 1; i <= 100; i++) {}`, ``)

	verify(t, `
	var s = 0
	for (var i = 1; i <= 100; i++) {
		s += i
	}
	println(s)`, `
	5050`)

	verify(t, `
	var s1 = 0
	var s2 = 0
	for (var i = 1; i <= 100; ++i) {
		if (i % 2 == 0) {
			s1 = s1 + i
		} else {
			s2 = s2 + i
		}
	}
	println(s1, s2)`, `
	2550 2500`)

	verify(t, `
	var s = 0
	for (var i = 0; i < 1000; i = i + 1) {
		if (i % 6 == 1 && (i % 7 == 2 || i % 8 == 3)) {
			s = s + i
		}
	}
	println(s)`, `
	29441`)

	verify(t, `
	var s = 0
	for (var i = 0; i < 1000; ++i) {
		if (i % 6 == 1 && i % 7 == 2 || i % 8 == 3) {
			s += i
		}
	}
	println(s)`, `
	71357`)

	verify(t, `
	var s = 0
	for (var i = 0; i < 10000; i += 1) {
		if (i % 3242 == 837) {
			break
		}
		s += i
	}
	println(s)`, `
	349866`)

	verify(t, `
	for (var i = 0; i < 100; ++i) {
		for (var j = 0; j < 100; ++j) {
			if ((i * j) % 12 == 7 && (i * j) % 23 == 11) {
				println(i, j)
				break
			}
		}
	}`, getOutput(func() {
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if (i*j)%12 == 7 && (i*j)%23 == 11 {
					fmt.Println(i, j)
					break
				}
			}
		}
	}))

	verify(t, `
	var s = 0
	for (var i = 0; i < 100; i = i + 1){
		if (i % 6 == 4) {
			continue
		}
		s += i * i
	}
	println(s)`, `
	277694`)

	verify(t, `
	for (var i = 0; i < 100; ++i) {
		for (var j = 0; j < 100; ++j) {
			if ((i * j) % 12 == 7 && (i * j) % 23 == 11) {
				continue
			}
			println(i, j)
		}
	}`, getOutput(func() {
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if (i*j)%12 == 7 && (i*j)%23 == 11 {
					continue
				}
				fmt.Println(i, j)
			}
		}
	}))
}

func TestWhile(t *testing.T) {
	verify(t, `
	var s = 0
	var i = 1
	while (i <= 100) {
		s = s + i
		i++
	}
	println(s, i)`, `
	5050 101`)

	verify(t, `
	var s1 = 0
	var s2 = 0
	var i = 1
	while (i <= 100) {
		if (i % 2 == 0) {
			s1 += i
		} else {
			s2 += i
		}
		i = i + 1
	}
	println(s1, s2, i)`, `
	2550 2500 101`)

	verify(t, `
	var s = 0
	var i = 0
	while (i < 10000) {
		if (i % 3242 == 837) {
			break
		}
		s = s + i
		i = i + 1
	}
	println(s, i)`, `
	349866 837`)

	verify(t, `
	var i = 0
	while (i < 100) {
		var j = 0
		while (j < 100) {
			if ((i * j) % 12 == 7 && (i * j) % 23 == 11) {
				println(i, j)
				break
			}
			j++
		}
		i++
	}`, getOutput(func() {
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if (i*j)%12 == 7 && (i*j)%23 == 11 {
					fmt.Println(i, j)
					break
				}
			}
		}
	}))

	verify(t, `
	var s = 0
	var i = 0
	while (i < 100) {
		if (i % 6 == 4) {
			++i
			continue
		}
		s = s + i * i
		++i
	}
	println(s, i)`, `
	277694 100`)

	verify(t, `
	var i = 0
	while (i < 100) {
		var j = 0
		while (j < 100) {
			if ((i * j) % 12 == 7 && (i * j) % 23 == 11) {
				j++
				continue
			}
			println(i, j)
			j++
		}
		i++
	}`, getOutput(func() {
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if (i*j)%12 == 7 && (i*j)%23 == 11 {
					continue
				}
				fmt.Println(i, j)
			}
		}
	}))
}

func TestClosure(t *testing.T) {
	verify(t, `
	var add = a => b => a + b
	println(add(2)(3))
	println(add(45)(67))
	var add5 = add(5)
	println(add5(7))
	println(add5(100))`, `
	5
	112
	12
	105`)

	verify(t, `
	var x = 100
	var fun = () => {x = x + 1}
	fun()
	fun()
	println(x)`, `
	102`)

	verify(t, `
	var x = 1000;
	(() => {x += 2})()
	println(x)`, `
	1002`)

	verify(t, `
	var compose = (n, f, g) => g(f(n))
	var f1 = n => n * 2
	var f2 = n => n + 1
	println(compose(100, f1, f2))`, `
	201`)

	verify(t, `
	var x = 123
	var outer = () => {
		var x = 456
		return () => x
	}
	println(x)
	println(outer()())
	println(x)`, `
	123
	456
	123`)

	verify(t, `
	var x = 123
	var outer = () => {
		x = 456
		return () => x
	}
	println(x)
	println(outer()())
	println(x)`, `
	123
	456
	456`)

	verify(t, `
	var x = 123
	var outer = () => {
		x = 456
		return () => x
	}
	x = 789
	println(x)
	println(outer()())
	println(x)`, `
	789
	456
	456`)

	verify(t, `
	var observer = callback => {
		for (var i = 1; i <= 10; i = i + 1) {
			callback(i)
		}
	}
	var s = 0
	observer(n => {s = s + n})
	println(s)`, `
	55`)

	verify(t, `
	var observer = callback => {
		for (var i = 1; i <= 10; i = i + 1) {
			callback(i)
		}
	}
	var s = 0
	observer(() => {s += 1})
	println(s)`, `
	10`)

	verify(t, `
	var Student = (name, age, score) => {
		return {
			getName: () => name,
			setName: _name => {name = _name},
			getAge() {
				return name			
			},
			setAge(_age) {
				age = _age
			},
			getScore: () => score,
			setScore: _score => {score = _score},
			getDescription: () => '(' + name + ' ' + age + ' ' + score + ')'
		}
	}
					
	var s1 = Student('Zhang San', 21, 87.5)
	var s2 = Student('Li Si', 23, 95)
	println(s1.getName())
	println(s2.getScore())
	println(s1.getDescription())
	println(s2.getDescription())
	s1.setName('Xiao Ming')
	s2.setScore(77.5)
	println(s1.getName())
	println(s2.getScore())
	println(s1.getDescription())
	println(s2.getDescription())`, `
	Zhang San
	95
	(Zhang San 21 87.5)
	(Li Si 23 95)
	Xiao Ming
	77.5
	(Xiao Ming 21 87.5)
	(Li Si 23 77.5)`)
}

func TestAssign(t *testing.T) {
	verify(t, `
	var arr = [1, 2, 3, [4, 5], {a: 100, b: 200, c: [100, 200]}]
	arr[0] = 5
	arr[1] += 10
	arr[3][1] *= 100
	arr[4].c[0] = 12
	arr[4].c[1] -= 8
					
	println(arr[0])
	println(arr[1])
	println(arr[3][1])
	println(arr[4].c[0])
	println(arr[4].c[1])`, `
	5
	12
	500
	12
	192`)

	verify(t, `
	var obj = {a: 100, b: 200, c: {d: 300, e: [{m: 10}, 2, 3]}}
	obj.a = 101
	obj.b -= 50
	obj.c.d += 100
	obj.c.e[0].m *= 10
					
	println(obj.a, obj.b, obj.c.d, obj.c.e[0].m)`, `
	101 150 400 100`)
}

func TestCallableLiteral(t *testing.T) {
	verify(t, `
	var f1 = () => 123
	println(f1())
	var f2 = () => 12.34
	println(f2())
	var f3 = () => 'hello'
	println(f3())
	var f4 = () => [100, 200, 300]
	println(f4())
	var f5 = () => (1 + 2) * 3
	println(f5())
	var f6 = () => 456
	var f7 = () => f6()
	println(f7())
	var f8 = () => {}
	println(f8())`, `
	123
	12.34
	hello
	[100, 200, 300]
	9
	456
	undefined`)

	verify(t, `
	var f1 = a => a + 1
	println(f1(10))
	var f2 = (a) => a + 1
	println(f2(20))
	var f3 = (a, b) => a + b
	println(f3(3, 5))
	var f4 = (a, b, c) => a * b * c
	println(f4(2, 4, 6))`, `
	11
	21
	8
	48`)

	verify(t, `
	var f1 = () => {return 100}
	println(f1())
	var f2 = () => {
		var a = 10
		var b = 20
		return a + b
	}
	println(f2())
	var f3 = () => {
		println('hello')
	}
	println(f3())
	
	var x = 1000
	var f4 = () => {x += 1}
	println(f4())
	println(x)`, `
	100
	30
	hello
	undefined
	undefined
	1001`)

	verify(t, `
	println((() => 12345)())
	println((m => m + 6)(10))
	println(((a, b) => a - b)(13, 7))
	
	var x = 10;
	((m, n) => {x += m + n})(12, 13)
	println(x)`, `
	12345
	16
	6
	35`)
}

func TestFunctionCall(t *testing.T) {
	verify(t, `
	var f1 = (a, b) => 123 * 456
	println(f1())
					
	var f2 = (a, b) => a + b
	println(f2(100, 200, 400))`, `
	56088
	300`)
}

func TestStringConcat(t *testing.T) {
	verify(t, `
	var s = ''
	for (var i = 1; i <= 10; i = i + 1) {
		if (i != 10) {
			s = s + i + ' '
		} else {
			s = s + i
		}
	}
	println(s)`, `
	1 2 3 4 5 6 7 8 9 10`)

	verify(t, `
	var s = ''
	for (var i = 0; i < 100; i = i + 1) {
		s += 'hello'
	}
	println(s)`, strings.Repeat("hello", 100))
}

func TestHarmonicSeries(t *testing.T) {
	verify(t, `
	var s = 0.0
	for (var i = 1; i <= 100; i++) {
		s += 1.0/i
	}
	println(s)`, getOutput(func() {
		s := 0.0
		for i := 1; i <= 100; i++ {
			s += 1.0 / float64(i)
		}
		fmt.Println(s)
	}))
}

func TestFibonacci(t *testing.T) {
	verify(t, `
	var fib = n => {
		if (n == 1) {
			return 1	
		} else if (n == 2) {
			return 1	
		}
		return fib(n - 1) + fib(n - 2)
	}
	println(fib(10))
	println(fib(20))
	`, `
	55
	6765`)

	verify(t, `
	function fib(n) {
		if (n == 1) {
			return 1	
		} else if (n == 2) {
			return 1	
		}
		return fib(n - 1) + fib(n - 2)
	}
	println(fib(10))
	println(fib(20))
	`, `
	55
	6765`)
}

func TestFactorial(t *testing.T) {
	verify(t, `
	var factorial = n => {
		if (n == 1) {
			return 1	
		}
		return n * factorial(n - 1)
	}

	println(factorial(10))
	`, `
	3628800`)

	verify(t, `
	function factorial(n) {
		if (n == 1) {
			return 1	
		}
		return n * factorial(n - 1)
	}

	println(factorial(10))
	`, `
	3628800`)
}

func TestCounter(t *testing.T) {
	verify(t, `
	// 计数器
	function Counter(init) {
		var cnt = init
		return {
			// 获取当前计数值
			current: () => cnt,
			// 计数值+1
			inc: () => {cnt++},
			// 计数值-1
			dec: () => {cnt--}
		}
	}

	var c1 = Counter(100)
	println(c1.current()) // 100
	c1.inc()
	println(c1.current()) // 101
	c1.inc()
	println(c1.current()) // 102
	c1.dec()
	println(c1.current()) // 101
	c1.dec()
	println(c1.current()) // 100
	
	var c2 = Counter(200)
	println(c2.current()) // 200
	c2.inc()
	println(c2.current()) // 201
	c2.inc()
	println(c2.current()) // 202
	c2.dec()
	println(c2.current()) // 201
	c2.dec()
	println(c2.current()) // 200
	`, `
	100
	101
	102
	101
	100
	200
	201
	202
	201
	200`)
}

func TestImport(t *testing.T) {
	importAndVerify(t, []string{"p1", "p2"}, `
	import a
	import b

	println('main')`, `
	d
	c
	b
	a
	main`)

	assert.Panics(t, func() {
		RunScript(`
		import x
                
		println('main')`, RunConfig{ImportPaths: []string{"p3"}})
	})
}

func TestTry(t *testing.T) {
	verify(t, `
	function testException(f) {
		try {
			f()
		} catch (e) {
			println('catch', e)
		} finally {
			println('finally')
		}
	}

	testException(() => {
		println('test1')
	})

	testException(() => {
		println('test2')
		throw 123
	})

	testException(() => {
		throw 456
		println('test3')
	})
	`, `
	test1
	finally
	test2
	catch 123
	finally
	catch 456
	finally`)

	verify(t, `
	try {
		println(123)
		throw 'hello'
		println(456)
	} catch (err) {
		println('catch', err)
	}
	`, `
	123
	catch hello`)

	verify(t, `
	try {
		println(123)
		throw 'hello'
		println(456)
	} catch (e) {
		println('catch1', e)
		try {
			throw 'hi'
		} catch (e) {
			println('catch2', e)
		}
	}
	`, `
	123
	catch1 hello
	catch2 hi`)

	verify(t, `
	function test() {
		println('test')
		try {
			println('try')
			return
		} catch (e) {
			println('catch')
		} finally {
			println('finally')
		}
	}
	test()
	`, `
	test
	try
	finally`)

	verify(t, `
	function test() {
		try {
			return 'try'
		} catch (e) {
			return 'catch'
		} finally {
			return 'finally'
		}
	}
	println(test())
	`, `
	finally`)
}
