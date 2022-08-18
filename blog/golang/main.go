/*
 * @Author: GG
 * @Date: 2022-08-17 15:13:06git
 * @LastEditTime: 2022-08-18 09:56:28
 * @LastEditors: GG
 * @Description: 原生golang 简单博客
 * @FilePath: \golang-demo\blog\golang\main.go
 *
 */
package main

import (
	"golang-demo/blog/golang/router"
	"log"
	"net/http"
)

func main() {
	// 程序入口，一个程序只有一个入口文件
	// web服务，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	router.Router()

	// server.ListenAndServe 监听端口并启动服务
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
