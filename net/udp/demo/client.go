/*
 * @Author: GG
 * @Date: 2023-04-26 16:37:36
 * @LastEditTime: 2023-04-26 16:40:59
 * @LastEditors: GG
 * @Description:
 * @FilePath: \demo\client.go
 *
 */
package main

import (
	"fmt"
	"net"
	"time"
)

var (
	serverAddr = "localhost:8080"
)

func main() {
	// 解析udp地址
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 建立udp连接
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 关闭udp连接
	defer conn.Close()

	// 发送消息
	msg := []byte("Hello, server!")
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取消息
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received:", string(buf[:n]))
}
