package main

import (
	"fmt"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
	"unsafe"
)

func main() {
	// pprof mem
	//Demo()
	Demo1()
	// wait 10 s
	runtime.GC()
	runtime.GC()
	time.Sleep(1 * time.Second)
	exit()
}
func exit() {
	fmt.Println("Saving heap profile...")
	// 打开一个文件用于保存 heap profile
	f, err := os.Create("heap.pprof")
	if err != nil {
		log.Fatal("could not create heap profile: ", err)
	}
	defer f.Close()
	runtime.GC()
	//time.Sleep(1 * time.Second)
	// 采集当前的 heap profile
	if err = pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write heap profile: ", err)
	}
	fmt.Println("Heap profile saved in heap.pprof")
	fmt.Println(packageStr0, packageStr1)
}

var packageStr0 []string // 一个包级变量
var packageStr1 []string

// 目前，s0和s1共享着承载它们的字节序列的同一个内存块。
// 虽然s1到这里已经不再被使用了，但是s0仍然在使用中，
// 所以它们共享的内存块将不会被回收。虽然此内存块中
// 只有50字节被真正使用，而其它字节却无法再被使用。
func Demo() {
	for i := 0; i < 10; i++ {
		s := createStringWithLengthOnHeap(1 << 20) //1M
		packageStr0 = append(packageStr0, s[:50])
	}

}
func Demo1() {
	for i := 0; i < 10; i++ {
		s := createStringWithLengthOnHeap(1 << 20) //1M
		packageStr1 = append(packageStr1, strings.Clone(s[:50]))
	}
}

func createStringWithLengthOnHeap(i int) string {
	s := make([]byte, i)
	for j := 0; j < i; j++ {
		s[j] = byte(j % 256 % (rand.Intn(256) + 1))
	}
	return string(s)
}

func toK8sBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
