package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close connection failed, err: ", err)
		}
	}(conn) // close connection before return
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // read
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("recv：", recvStr)
		if recvStr == "exit\r\n" {
			fmt.Println("client request to exit")
			break
		}
		_, err = conn.Write([]byte(strings.ToUpper("Hello " + recvStr)))
		if err != nil {
			return
		} // send
	}
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept() // listen
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		go process(conn) // handle connection concurrently
	}
}
