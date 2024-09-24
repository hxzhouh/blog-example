package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	port       string
	bufferSize int64
)

func init() {
	flag.StringVar(&port, "p", "9000", "Port to run the TCP server on")
	flag.Int64Var(&bufferSize, "b", 1024_000, "Buffer size for reading data")
	flag.Parse()
}
func main() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	file, err := os.Create("/dev/null")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Failed to read from connection: %v\n", err)
			}
			break
		}
		if _, err := file.Write(buffer[:n]); err != nil {
			fmt.Printf("Failed to write to file: %v\n", err)
			break
		}
	}
	fmt.Println("File received successfully")
}
