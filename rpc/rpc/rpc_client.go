/*
 * @Author: GG
 * @Date: 2022-10-06 15:44:43
 * @LastEditTime: 2022-10-06 15:54:14
 * @LastEditors: GG
 * @Description:
 * @FilePath: \rpc\client\rpc_client.go
 *
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 请求结构体
type ArithRequest struct {
	A int
	B int
}

// 响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {

	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("dailing error:", err)
	}

	req := ArithRequest{9, 2}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
