/*
 * @Author: GG
 * @Date: 2023-04-26 16:37:47
 * @LastEditTime: 2023-04-26 17:19:22
 * @LastEditors: GG
 * @Description:
 * @FilePath: \demo\server.go
 *
 */
package main

import (
	"fmt"
	"net"
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

	// 建立udp监听
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// 接收请求
		buf := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Received %d bytes from %s:%s\n", n, clientAddr, string(buf[:n]))

		// 发送响应
		_, err = conn.WriteToUDP(buf[:n], clientAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Send %d bytes to %s:\n", n, clientAddr)
	}
}
