package main

import (
	"blog-example/hrpc/api"
	"log"
	"testing"
)

func TestSppRpcCall(t *testing.T) {

	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	var reply api.String

	err = client.Hello(api.String{Value: "hello"}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}
