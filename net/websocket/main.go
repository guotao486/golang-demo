/*
 * @Author: GG
 * @Date: 2022-11-28 18:20:07
 * @LastEditTime: 2022-11-28 18:31:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \net\websocket\main.go
 *
 */
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	// 设定websocket服务信息处理
	http.Handle("/", websocket.Handler(server))

	// 监听
	http.ListenAndServe(":7777", nil)
}

func server(ws *websocket.Conn) {
	fmt.Println("new connection")

	data := make([]byte, 1024)
	for {
		// 读取信息
		d, err := ws.Read(data)
		if err != nil {
			fmt.Println("err ,", err)
			break
		}
		fmt.Println("读取的信息：", d)
		ws.Write([]byte("hello i'm is webscoket server"))
	}
}
