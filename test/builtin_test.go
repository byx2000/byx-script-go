package test

import "testing"

func TestIsType(t *testing.T) {
	verify(t, `
	var objs = [123, 3.14, true, 'hello', [1, 2, 3], {a: 100, b: 'hi'}, undefined]
                
	for (var i = 0; i < objs.length(); ++i) {
		print(isInteger(objs[i]) + ' ')
	}
	println()
	
	for (var i = 0; i < objs.length(); ++i) {
		print(isDouble(objs[i]) + ' ')
	}
	println()
	
	for (var i = 0; i < objs.length(); ++i) {
		print(isBool(objs[i]) + ' ')
	}
	println()
	
	for (var i = 0; i < objs.length(); ++i) {
		print(isString(objs[i]) + ' ')
	}
	println()
	
	for (var i = 0; i < objs.length(); ++i) {
		print(isList(objs[i]) + ' ')
	}
	println()
	
	for (var i = 0; i < objs.length(); ++i) {
		print(isObject(objs[i]) + ' ')
	}
	println()
	
	for (var i = 0; i < objs.length(); ++i) {
		print(isUndefined(objs[i]) + ' ')
	}
	println()
	`, `
	true false false false false false false
	false true false false false false false
	false false true false false false false
	false false false true false false false
	false false false false true false false
	false false false false false true false
	false false false false false false true`)
}

func TestHashcode(t *testing.T) {
	verify(t, `
	var objs = [123, 3.14, true, 'hello', [1, 2, 3], {a: 100, b: 'hi'}, undefined]
	for (var i = 0; i < objs.length(); ++i) {
		for (var j = 0; j < objs.length(); ++j) {
			print((hashcode(objs[i]) == hashcode(objs[j])) + ' ')
		}
		println()
	}
	`, `
	true false false false false false false
	false true false false false false false
	false false true false false false false
	false false false true false false false
	false false false false true false false
	false false false false false true false
	false false false false false false true`)
}
