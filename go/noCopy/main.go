package main

import "fmt"

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type StructA struct {
	Name string
}
type StructB struct {
	noCopy
	Age  int
	Name *StructA
}

func main() {
	a := StructB{}
	a.Age = 30
	a.Name = &StructA{"Tom"}
	b := a
	b.Age = 20
	b.Name.Name = "Jerry"
	fmt.Println(a.Name, b.Name)
}
