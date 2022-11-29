/*
 * @Author: GG
 * @Date: 2022-11-28 16:05:32
 * @LastEditTime: 2022-11-29 16:42:40
 * @LastEditors: GG
 * @Description: net tcp
 * @FilePath: \net\tcp\server\main.go
 *
 */
package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
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
	reader := bufio.NewReader(c)
	for {
		// var data [1024]byte

		// n, err := bufio.NewReader(c).Read(data[:])
		// n, err := c.Read(data[:])
		msg, err := unpack(reader)
		if err != nil {
			fmt.Printf("reader err: %v\n", err)
			return
		}

		fmt.Println(msg)
		c.Write([]byte("hello worlds!"))
	}
}

func unpack(reader *bufio.Reader) (string, error) {
	lenByte, _ := reader.Peek(2) // 读取前 固定的几个字节信息
	lengthBuff := bytes.NewBuffer(lenByte)
	var length int16
	err := binary.Read(lengthBuff, binary.BigEndian, &length)
	fmt.Printf("length: %v\n", length)
	if err != nil {
		return "", err
	}

	// 获取信息
	// 包头 + 数据长度 = length
	if int16(reader.Buffered()) < length+2 {
		return "", err
	}

	// 真正的读取
	pack := make([]byte, int(2+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[2:]), nil
}
