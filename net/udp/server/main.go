/*
 * @Author: GG
 * @Date: 2022-11-28 16:56:57
 * @LastEditTime: 2022-11-28 17:36:44
 * @LastEditors: GG
 * @Description:
 * @FilePath: \net\udp\server\main.go
 *
 */
package main

import (
	"fmt"
	"net"
)

func main() {

	listen, _ := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 5555,
	})

	defer listen.Close()

	for {
		// udp 直接接收信息
		var data [1024]byte

		n, addr, err := listen.ReadFromUDP(data[:]) //接收信息
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		fmt.Printf("data: %v,addr:%s\n", string(data[:n]), addr)

		// 发送信息
		listen.WriteToUDP([]byte("i'm is udp"), addr)
	}
}
