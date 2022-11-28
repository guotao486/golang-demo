/*
 * @Author: GG
 * @Date: 2022-11-28 16:05:25
 * @LastEditTime: 2022-11-28 16:45:08
 * @LastEditors: GG
 * @Description:
 * @FilePath: \net\tcp\client\main.go
 *
 */
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("连接tcp客户端，127.0.0.1:3333")
	// 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 先写入
	conn.Write([]byte("hello server"))

	// 读取
	var data [1024]byte
	n, err := bufio.NewReader(conn).Read(data[:])
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("n: %v\n", string(data[:n]))

}
