package main

import (
	"fmt"
	"net"
	"strings"
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

	// 建立连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	defer conn.Close()

	fmt.Println("Server started on ", serverAddr)

	clients := make(map[string]*net.UDPAddr)

	for {
		// 接收消息
		buf := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		// 解析消息
		msg := strings.TrimSpace(string(buf[:n]))
		if msg == "" {
			continue
		}

		fmt.Printf("Received %d bytes from %s:%s\n", n, clientAddr, msg)

		if _, ok := clients[clientAddr.String()]; !ok {
			fmt.Println("New Client:", clientAddr.String())
			clients[clientAddr.String()] = clientAddr
		}

		for _, addr := range clients {
			if addr.String() == clientAddr.String() {
				continue
			}

			_, err := conn.WriteToUDP([]byte(msg), addr)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				continue
			}

			fmt.Printf("Send %d bytes to %s: %s\n", len(msg), addr, msg)
		}
	}
}
