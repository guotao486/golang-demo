package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头，允许跨域请求
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // 允许所有域名访问
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // 允许的HTTP方法
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // 允许的HTTP请求头

		// 处理跨域请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK) // 返回200状态码表示预检请求成功
			return
		}

		// 返回内容
		fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":8080", nil)
}
