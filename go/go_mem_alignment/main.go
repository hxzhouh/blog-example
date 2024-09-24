package main

import (
	"fmt"
	"unsafe"
)

type T1 struct {
	a int8
	b int64
	c int16
}

type T2 struct {
	a int8
	c int16
	b int64
}

func main() {
	t := T1{}
	fmt.Println(fmt.Sprintf("%d %d %d %d", unsafe.Sizeof(t.a), unsafe.Sizeof(t.b), unsafe.Sizeof(t.c), unsafe.Sizeof(t)))
	fmt.Println(fmt.Sprintf("%p %p %p", &t.a, &t.b, &t.c))
	fmt.Println(unsafe.Alignof(t))
}
