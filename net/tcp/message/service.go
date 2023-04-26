/*
 * @Author: GG
 * @Date: 2023-04-26 14:37:49
 * @LastEditTime: 2023-04-26 16:13:46
 * @LastEditors: GG
 * @Description:
 * @FilePath: \message\service.go
 *
 */
package main

import (
	"fmt"
	"net"
)

var (
	clients    = make(map[net.Addr]net.Conn) // 客户端
	addCh      = make(chan net.Conn)         // 新增链接
	delCh      = make(chan net.Addr)         // 删除链接
	messageCh  = make(chan []byte)           // 消息
	listenAddr = "127.0.0.1:8080"
)

func main() {
	fmt.Println("Service started on", listenAddr)
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listen.Close()

	go broadcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 获取一个新连接
		addCh <- conn
		go handleRequest(conn)
	}

}

func broadcaster() {
	select {
	case conn := <-addCh:
		clients[conn.RemoteAddr()] = conn
		fmt.Println("New Client:", conn.RemoteAddr())
	case addr := <-delCh:
		delete(clients, addr)
		fmt.Println("Client disconnected:", addr)
	case msg := <-messageCh:
		for _, client := range clients {
			_, err := client.Write(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// 处理请求
func handleRequest(conn net.Conn) {
	defer func() {
		delCh <- conn.RemoteAddr()
		conn.Close()
	}()

	for {
		msg := make([]byte, 4096)
		n, err := conn.Read(msg)
		if err != nil {
			return
		}

		messageCh <- msg[:n]

	}
}
