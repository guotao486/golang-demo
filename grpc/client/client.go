/*
 * @Author: GG
 * @Date: 2022-10-07 15:33:01
 * @LastEditTime: 2022-10-07 17:29:19
 * @LastEditors: GG
 * @Description:
 * @FilePath: \grpc\client\client.go
 *
 */
package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		fmt.Printf("os.Args: %v\n", os.Args)
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	// StreamClient(c)
	// Streamserver(c)
	StreamAction(c)
}

// 接收流
func StreamClient(c pb.GreeterClient) {
	//调用服务端推送流
	reqstreamData := &pb.StreamReqData{Data: "aaa"}
	res, _ := c.StreamClient(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
}

// 发送流
func Streamserver(c pb.GreeterClient) {
	//客户端 推送 流
	putRes, _ := c.Streamserver(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&pb.StreamReqData{Data: "ss"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}

func StreamAction(c pb.GreeterClient) {
	//服务端 客户端 双向流
	allStr, _ := c.StreamAction(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pb.StreamReqData{Data: "from client"})
			time.Sleep(time.Second)
		}
	}()
	select {}
}
