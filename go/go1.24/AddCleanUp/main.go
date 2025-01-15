package main

import (
	"fmt"
	"runtime"
	"time"
)

type MyStruct struct {
	Name  string
	Other *MyStruct
}

func main() {
	x := MyStruct{Name: "X"}
	y := MyStruct{Name: "Y"}

	x.Other = &y
	y.Other = &x
	//runtime.SetFinalizer(&x, func(x *MyStruct) {
	//	fmt.Printf("Finalizer for %s is called\n", x.Name)
	//})
	//runtime.SetFinalizer(&y, func(y *MyStruct) {
	//	fmt.Printf("Finalizer for %s is called\n", y.Name)
	//})
	xName := x.Name
	runtime.AddCleanup(&x, func(name string) {
		fmt.Println("Cleanup for", x)
	}, xName)
	yName := y.Name
	runtime.AddCleanup(&y, func(name string) {
		fmt.Println("Cleanup for", x)
	}, yName)
	time.Sleep(time.Millisecond)
	runtime.GC()
	time.Sleep(time.Millisecond)
	runtime.GC()
}
