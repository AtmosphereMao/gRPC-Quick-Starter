package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
	"zyy-grpc-server/src/service"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewHelloServiceClient(conn)

	// 客户端先调用Channel方法，获取返回的流对象

	stream, err01 := client.Channel(context.Background())
	if err01 != nil {
		log.Fatal(err01)
	}

	// 在客户端我们将发送和接收放到两个独立的Goroutine

	// 首先向服务端发送数据：
	go func() {
		for {
			req := &service.Request{
				Value: "test",
			}
			if err := stream.Send(req); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 然后再循环中接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
