/*
 * @Author: GG
 * @Date: 2022-11-28 16:05:32
 * @LastEditTime: 2022-11-28 16:54:19
 * @LastEditors: GG
 * @Description: net tcp
 * @FilePath: \net\tcp\server\main.go
 *
 */
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("启动tcp客户端，127.0.0.1:3333")
	// 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Printf("listen err: %v\n", err)
		return //表示程序结束
	}

	for {
		// 接受客户端建立的链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept err: %v\n", err)
			break
		}

		// 处理用户连接信息
		go handler(conn)
	}
}

func handler(c net.Conn) {

	defer c.Close()
	for {
		var data [1024]byte

		n, err := bufio.NewReader(c).Read(data[:])
		if err != nil {
			fmt.Printf("reader err: %v\n", err)
			break
		}

		fmt.Printf("n: %v\n", string(data[:n]))
		c.Write([]byte("hello worlds!"))
	}
}
