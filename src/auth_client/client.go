package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"zyy-grpc-server/src/author"
	"zyy-grpc-server/src/service"
)

func main() {
	conn, err := grpc.Dial("localhost:8888",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(author.NewClientAuthentication("admin", "123456")),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// NewHelloServiceClient函数是xxx_grpc.pb.go中自动生成的函数，
	// 基于已经建立的连接构造HelloServiceClient对象,
	// 返回的client其实是一个HelloServiceClient接口对象
	//
	client := service.NewHelloServiceClient(conn)

	// 通过接口定义的方法就可以调用服务端对应gRPC服务提供的方法
	req := &service.Request{Value: "小亮"}
	reply, err01 := client.Hello(context.Background(), req)
	if err01 != nil {
		log.Fatal(err01)
	}
	fmt.Println(reply.GetValue())
}
