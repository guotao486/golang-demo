/*
 * @Author: GG
 * @Date: 2022-10-06 14:20:01
 * @LastEditTime: 2022-10-06 16:14:07
 * @LastEditors: GG
 * @Description:rpc 服务端，只适用于golang
 * @FilePath: \rpc\jsonrpc_service.go
 *
 */
package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

// 乘法运算
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}

	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

// net/rpc
func main() {
	rpc.Register(new(Arith)) // 注册rpc 服务

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection")

	for {
		conn, err := lis.Accept() // 接收客户端连接
		if err != nil {
			continue
		}

		// 处理并发消息
		go func(c net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
