package main

import (
	"testing"
	"weak"
)

//type User struct {
//	Name string
//	Age  int
//	Properties
//}
//type Properties struct {
//	Height int
//	Weight int
//}

var pointMap map[int]*int
var noPointMap map[int]int
var weakMap map[weak.Pointer[int]]int

func BenchmarkPointMap(b *testing.B) {
	pointMap = make(map[int]*int)
	for i := 0; i < 10e6; i++ {
		pointMap[i] = &i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(pointMap, i)
		pointMap[i] = &i
	}
}

func BenchmarkNoPointMap(b *testing.B) {

	noPointMap = make(map[int]int)
	for i := 0; i < 10e6; i++ {
		noPointMap[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(noPointMap, i)
		noPointMap[i] = i
	}
}

func BenchmarkWeakMap(b *testing.B) {
	weakMap = make(map[weak.Pointer[int]]int)
	for i := 0; i < 10e6; i++ {
		kw := weak.Make(&i)
		noPointMap[kw] = i
	}
}
