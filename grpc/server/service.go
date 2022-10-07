/*
 * @Author: GG
 * @Date: 2022-10-07 14:55:35
 * @LastEditTime: 2022-10-07 17:24:49
 * @LastEditors: GG
 * @Description:
 * @FilePath: \grpc\server\service.go
 *
 */
package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "127.0.0.1:50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// 普通接口
func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil
}

// 服务端接收流，客户端发送
func (s server) Streamserver(cliStr pb.Greeter_StreamserverServer) error {
	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem.Data)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

// 服务端发送流，客户端接收
func (s server) StreamClient(req *pb.StreamReqData, res pb.Greeter_StreamClientServer) error {
	i := 0
	for {
		i++
		res.Send(&pb.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 双向流
func (s server) StreamAction(allStr pb.Greeter_StreamActionServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pb.StreamResData{Data: "from server"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 启动grpc服务
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
