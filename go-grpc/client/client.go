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
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8098", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()

	// 新建一个客户端
	c := hello_grpc.NewGreeterClient(conn)

	// 调用服务端函数
	r, err := c.SayHello(context.Background(), &hello_grpc.HelloRequest{Name: "horika"})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Printf("调用成功: %s", r.Message)
}
