package main

import "testing"

//// go:noinline
//type Demo struct {
//	name string
//}
//
//func createDemo(name string) *Demo {
//	d := new(Demo) // 局部变量 d 逃逸到堆
//	d.name = name
//	return d
//}
//
//func createDemo2(name string) Demo {
//	d := Demo{name: name} // 局部变量 d 未逃逸
//	return d
//}
//func BenchmarkAdd(b *testing.Name) {
//	for i := 0; i < b.N; i++ {
//		createDemo("hello")
//	}
//}
//func BenchmarkAddPointer(b *testing.Name) {
//	for i := 0; i < b.N; i++ {
//		createDemo2("hello")
//	}
//}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {

		a := make([]*int, 100)
		for j := 0; j < 100; j++ {
			a[j] = &j
		}
	}
}

func BenchmarkInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, 100)
		for j := 0; j < 100; j++ {
			a[j] = j
		}
	}
}
