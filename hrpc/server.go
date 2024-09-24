package main

import (
	"blog-example/hrpc/api"
	"log"
	"net/rpc"
)

const ServerName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request api.String, reply *api.String) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(ServerName, srv)
}

type HelloService struct{}

func (p *HelloService) Hello(request api.String, reply *api.String) error {
	log.Println("HelloService.proto Hello")
	reply.Value = "hello:" + request.GetValue()
	return nil
}
