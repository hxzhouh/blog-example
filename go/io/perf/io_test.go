package main

import (
	"io"
	"os"
	"testing"
)

// testFile
// mac: mkfile test_file_100M.txt
// linux: dd if=/dev/zero of=test_file_100M bs=1m count=100
const filePath = "test_file_4k.txt"
const filePath_100M = "test_file_100M.txt"
const filePath_1b = "test_file_1b.txt"

// Normal IO read file
func normalRead(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// read all
	_, err = io.ReadAll(file)
	// receive data
	tempFile, err := os.Create("/dev/null")
	if err != nil {
		return err
	}
	io.Copy(tempFile, file)
	return err
}

// Streaming IO reads files
func streamRead(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// receive data
	tempFile, err := os.Create("/dev/null")
	if err != nil {
		return err
	}
	defer func() {
		_ = tempFile.Close()
	}()
	_, err = tempFile.ReadFrom(file)
	return err
}

func BenchmarkNormalRead100M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = normalRead(filePath_100M)
	}
}

func BenchmarkStreamRead100M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = streamRead(filePath_100M)
	}
}

func BenchmarkNormalRead4k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = normalRead(filePath)
	}
}

func BenchmarkStreamRead4k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = streamRead(filePath)
	}
}

// read
func BenchmarkNormalRead1b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = normalRead(filePath_1b)
	}
}

// streaming reading
func BenchmarkStreamRead1b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = streamRead(filePath_1b)
	}
}

//// Benchmark 普通写入
//func BenchmarkNormalWrite(b *testing.B) {
//	data := []byte("some test data")
//	for i := 0; i < b.N; i++ {
//		_ = normalWrite(filePath, data)
//	}
//}
//
//// Benchmark 流式写入
//func BenchmarkStreamWrite(b *testing.B) {
//	data := []byte("some test data")
//	for i := 0; i < b.N; i++ {
//		_ = streamWrite(filePath, data)
//	}
//}
