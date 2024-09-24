package main

// MyStruct 是一个简单的结构体，包含一个指针字段。
type MyStruct struct {
	Name  string
	Other *MyStruct
}

//func main() {
//	x := MyStruct{Name: "X"}
//	y := MyStruct{Name: "Y"}
//
//	x.Other = &y
//	y.Other = &x
//	runtime.SetFinalizer(&x, func(x *MyStruct) {
//		fmt.Printf("Finalizer for %s is called\n", x.Name)
//	})
//	time.Sleep(time.Second)
//	runtime.GC()
//	time.Sleep(time.Second)
//}
