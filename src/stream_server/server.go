package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"zyy-grpc-server/src/service"
)

type HelloService struct {
	service.UnimplementedHelloServiceServer
}

var _ service.HelloServiceServer = (*HelloService)(nil)

func (p *HelloService) Channel(stream service.HelloService_ChannelServer) error {
	// 服务端在循环中接收客户端发来的数据
	for {
		args, err := stream.Recv()
		if err != nil {
			// 如果遇到io.EOF表示客户端流关闭
			if err == io.EOF {
				return nil
			}
			return err
		}

		// 响应一个请求
		// 生成返回的数据通过流发送给客户端
		resp := &service.Response{
			Value: "hello," + args.GetValue(),
		}
		log.Printf("Received: %v", resp.Value)
		err = stream.Send(resp)
		if err != nil {
			// 服务端发送异常，函数退出，服务端流关闭
			return err
		}
	}
}

func main() {
	// 1. 构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 2. 通过gRPC插件生成的RegisterHelloServiceServer 函数注册我们实现的HelloService服务。
	service.RegisterHelloServiceServer(grpcServer, new(HelloService))
	// 3. 监听:8888端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Listen TCP err:", err)
	}
	// 4. 通过grpcServer.Serve(listen) 在一个监听端口上提供gRPC服务
	grpcServer.Serve(listen)
}
