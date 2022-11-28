/*
 * @Author: GG
 * @Date: 2022-11-28 16:59:02
 * @LastEditTime: 2022-11-28 17:34:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \net\udp\client\main.go
 *
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	socket, _ := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 5555,
	})

	defer socket.Close()

	// 发送数据
	socket.Write([]byte("hello udp server"))

	// 接收数据
	var data [1024]byte
	n, addr, _ := socket.ReadFromUDP(data[:])
	fmt.Printf("data: %v,addr:%s\n", string(data[:n]), addr)
}
