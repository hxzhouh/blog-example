package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is running on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Client disconnected:", conn.RemoteAddr())
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from client: %v\n", err)
			return
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Received: %s\n", message)

		if strings.ToLower(message) == "quit" {
			conn.Write([]byte("Goodbye!\n"))
			return
		}

		conn.Write([]byte("Echo: " + message + "\n"))
	}
}
