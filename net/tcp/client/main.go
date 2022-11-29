/*
 * @Author: GG
 * @Date: 2022-11-28 16:05:25
 * @LastEditTime: 2022-11-29 16:45:55
 * @LastEditors: GG
 * @Description:
 * @FilePath: \net\tcp\client\main.go
 *
 */
package main

import (
	"bytes"
	"encoding/binary"
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

	for i := 0; i < 10; i++ {

		// tcp 粘包问题,包头2个字节
		// int16 占用 2个字节
		msg := "shineyork666"
		msgLen := len(msg)
		length := int16(msgLen) // 将信息长度转换成int16，长度2个字节

		pkg := new(bytes.Buffer)
		binary.Write(pkg, binary.BigEndian, length) // 信息长度 设置为包头
		data := append(pkg.Bytes(), []byte(msg)...)
		fmt.Printf("data: %v\n", string(data))

		// 先写入
		conn.Write(data)
	}
	// 读取
	var pack [1024]byte
	// n, err := bufio.NewReader(conn).Read(pack[:])
	n, err := conn.Read(pack[:])
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("n: %v\n", string(pack[:n]))
	conn.Close()
}
