package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		f, _ := os.Open("./testmmap.txt")
		fullCopy(f, writer)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func fullCopy(read *os.File, write http.ResponseWriter) {
	buf := make([]byte, 1024)
	n, _ := read.Read(buf)
	fmt.Println(fmt.Sprintf("%d", &buf))
	_, _ = write.Write(buf[:n])
}

func mmapCopy(read *os.File, writer http.ResponseWriter) {
	buf, err := syscall.Mmap(int(read.Fd()), 0, 11, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%d", &buf))
	_, _ = writer.Write(buf)
}
