package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const (
	serverAddress = "localhost:9000"
	bufferSize    = 1024 // Control the buffer size here
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run point_test.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, bufferSize)
	start := time.Now()
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Failed to read from file: %v\n", err)
			}
			break
		}
		if _, err := conn.Write(buffer[:n]); err != nil {
			fmt.Printf("Failed to write to connection: %v\n", err)
			break
		}
	}
	fmt.Println(fmt.Sprintf("File sent successfully cost :%d ms", time.Since(start).Milliseconds()))
}
