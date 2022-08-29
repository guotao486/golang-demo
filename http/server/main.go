/*
 * @Author: GG
 * @Date: 2022-08-29 17:31:01
 * @LastEditTime: 2022-08-29 17:34:52
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\http\server\main.go
 *
 */
package main

import (
	"io"
	"net/http"
)

// 方式1
func test1() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	// 请求处理函数
	f := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "hello world")
	}
	// 响应路径,注意前面要有斜杠 /
	http.HandleFunc("/hello", f)
	server.ListenAndServe()
}

//方式2
func test2() {
	// 请求处理函数
	f := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "hello world")
	}
	// 响应路径,注意前面要有斜杠 /
	http.HandleFunc("/hello", f)
	http.ListenAndServe(":8080", nil)
}
func main() {
	// test1()
	test2()
}
