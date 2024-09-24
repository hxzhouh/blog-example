package main

import (
	"blog-example/hrpc/api"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: client}, nil
}

func (p *HelloServiceClient) Hello(request api.String, reply *api.String) error {
	return p.Client.Call(ServerName+".Hello", request, reply)
}
