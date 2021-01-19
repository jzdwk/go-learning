/*
@Time : 20-10-22
@Author : jzd
@Project: go-learning
*/
package main

import (
	"context"
	"fmt"
	"go-learning/go-grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *hello_grpc.HelloRequest) (*hello_grpc.HelloReply, error) {
	return &hello_grpc.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":8098")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	hello_grpc.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
	fmt.Println("start gRpc...")
}
