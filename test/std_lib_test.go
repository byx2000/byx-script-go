package test

import (
	"fmt"
	"math"
	"testing"
)

func verifyWithLib(t *testing.T, script string, expectedOutput string) {
	importAndVerify(t, []string{"lib"}, script, expectedOutput)
}

func TestStack(t *testing.T) {
	verifyWithLib(t, `
	import stack
                
	var s = Stack()
	s.push(1)
	s.push(2)
	s.push(3)
	println(s.size(), s.isEmpty())
	
	println(s.pop())
	println(s.pop())
	println(s.size(), s.isEmpty())
	
	s.push(4)
	s.push(5)
	println(s.top())
	println(s.size())
	
	println(s.pop())
	println(s.pop())
	println(s.pop())
	println(s.size(), s.isEmpty())`, `
	3 false
	3
	2
	1 false
	5
	3
	5
	4
	1
	0 true`)
}

func TestQueue(t *testing.T) {
	verifyWithLib(t, `
	import queue
                
	var q = Queue()
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	println(q.size(), q.isEmpty())
	
	println(q.deQueue())
	println(q.deQueue())
	println(q.size(), q.isEmpty())
	
	q.enQueue(4)
	q.enQueue(5)
	q.enQueue(6)
	println(q.size())
	println(q.front())
	println(q.tail())
	
	println(q.deQueue())
	println(q.deQueue())
	println(q.deQueue())
	q.deQueue()
	println(q.size(), q.isEmpty())`, `
	3 false
	1
	2
	1 false
	4
	3
	6
	3
	4
	5
	0 true`)
}

func TestStream(t *testing.T) {
	verifyWithLib(t, `
	import stream
                
	Stream.of([1, 2, 3, 4, 5])
		.map(n => n + 1)
		.filter(n => n % 2 == 0)
		.forEach(println)`, `
	2
	4
	6`)

	verifyWithLib(t, `
	import stream
                
	var s = Stream.of([1, 2, 3, 4, 5])
		.map(n => n + 1)
		.filter(n => n % 2 == 1)
		.toList()
	println(s)
	`, `
	[3, 5]`)
}

func TestSet(t *testing.T) {
	verifyWithLib(t, `
	import set
                
	var set = Set()
	set.add(1)
	set.add(2)
	set.add(3)
	set.add(2)
	var list = set.toList()
	println(set.size())
	println(set.contains(1))
	println(set.contains(2))
	println(set.contains(3))
	println(set.contains(4))
	println(set.remove(1))
	println(set.remove(5))
	println(set.size())
	println(set.contains(1))
	println(set.contains(2))
	println(set.contains(3))
	println(set.contains(4))
	println(set.isEmpty())
	set.remove(2)
	set.remove(3)
	println(set.isEmpty())`, `
	3
	true
	true
	true
	false
	true
	false
	2
	false
	true
	true
	false
	false
	true`)

	verifyWithLib(t, `
	import set
                
	var s = Set()
	for (var i = 0; i < 100; ++i) {
		s.add(i % 10)
	}
	println(s.size())`, `
	10`)
}

func TestMap(t *testing.T) {
	verifyWithLib(t, `
	import map
                
	var map = Map()
	map.put('k1', 123)
	map.put('k2', 456)
	map.put('k3', 789)
	println(map.size())
	println(map.get('k1'))
	println(map.get('k2'))
	println(map.get('k3'))
	println(map.get('k4'))
	println(map.containsKey('k1'))
	println(map.containsKey('k2'))
	println(map.containsKey('k3'))
	println(map.containsKey('k4'))
	println(map.put('k2', 12345))
	println(map.get('k2'))
	println(map.remove('k4'))
	println(map.remove('k1'))
	println(map.size())`, `
	3
	123
	456
	789
	undefined
	true
	true
	true
	false
	456
	12345
	undefined
	123
	2`)

	verifyWithLib(t, `
	import map
                
	function twoSum(nums, target) {
		var map = Map()
		for (var i = 0; i < nums.length(); ++i) {
			if (map.containsKey(target - nums[i])) {
				return [map.get(target - nums[i]), i]
			}
			map.put(nums[i], i)
		}
		return undefined
	}
	
	println(twoSum([2, 7, 11, 15], 9))
	println(twoSum([3, 2, 4], 6))
	println(twoSum([3, 3], 6))
	println(twoSum([23, 16, 76, 97, 240, 224, 5, 78, 443, 25], 103))`, `
	[0, 1]
	[1, 2]
	[0, 1]
	[7, 9]`)
}

func TestMath(t *testing.T) {
	verifyWithLib(t, `
	import math
                
	println(Math.abs(15))
	println(Math.abs(-3.14))
	println(Math.sin(10))
	println(Math.sin(12.34))
	println(Math.cos(10))
	println(Math.cos(12.34))
	println(Math.tan(10))
	println(Math.tan(12.34))
	println(Math.pow(2, 3))
	println(Math.pow(2, 3.5))
	println(Math.pow(2.5, 3))
	println(Math.pow(2.5, 3.5))
	println(Math.exp(2))
	println(Math.exp(3.14))
	println(Math.ln(25))
	println(Math.ln(12.56))
	println(Math.log10(25))
	println(Math.log10(12.56))
	println(Math.sqrt(2))
	println(Math.sqrt(31.5))
	println(Math.round(7))
	println(Math.round(8.3))
	println(Math.round(12.9))
	println(Math.ceil(7))
	println(Math.ceil(8.3))
	println(Math.ceil(12.9))
	println(Math.floor(7))
	println(Math.floor(8.3))
	println(Math.floor(12.9))`, getOutput(func() {
		fmt.Println(math.Abs(15))
		fmt.Println(math.Abs(-3.14))
		fmt.Println(math.Sin(10))
		fmt.Println(math.Sin(12.34))
		fmt.Println(math.Cos(10))
		fmt.Println(math.Cos(12.34))
		fmt.Println(math.Tan(10))
		fmt.Println(math.Tan(12.34))
		fmt.Println(math.Pow(2, 3))
		fmt.Println(math.Pow(2, 3.5))
		fmt.Println(math.Pow(2.5, 3))
		fmt.Println(math.Pow(2.5, 3.5))
		fmt.Println(math.Exp(2))
		fmt.Println(math.Exp(3.14))
		fmt.Println(math.Log(25))
		fmt.Println(math.Log(12.56))
		fmt.Println(math.Log10(25))
		fmt.Println(math.Log10(12.56))
		fmt.Println(math.Sqrt(2))
		fmt.Println(math.Sqrt(31.5))
		fmt.Println(math.Round(7))
		fmt.Println(math.Round(8.3))
		fmt.Println(math.Round(12.9))
		fmt.Println(math.Ceil(7))
		fmt.Println(math.Ceil(8.3))
		fmt.Println(math.Ceil(12.9))
		fmt.Println(math.Floor(7))
		fmt.Println(math.Floor(8.3))
		fmt.Println(math.Floor(12.9))
	}))
}
