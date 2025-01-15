package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	root, err := os.OpenRoot(".")
	if err != nil {
		panic(err)
	}
	file, err := root.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %s: %s\n", fileName, err.Error()))
		return
	}
	content := make([]byte, 1024)
	c, err := file.Read(content)
	if err != nil {
		panic(err)
	}
	content = content[:c]
	fmt.Printf("File %s opened successfully. file content %s\n", fileName, content)
}
