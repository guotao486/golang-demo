/*
 * @Author: GG
 * @Date: 2023-04-28 10:19:41
 * @LastEditTime: 2023-04-28 11:18:01
 * @LastEditors: GG
 * @Description:
 * @FilePath: \http\demo\demo.go
 *
 */
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	server()
}

func get() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("body: %v\n", string(body))
}

func post() {

	url := "http://www.example.com/login"
	data := []byte(`{"username": "admin", "password": "password"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	// 静态文件
	http.Handle("/file", http.FileServer(http.Dir("/")))

	http.ListenAndServe(":8080", nil)
}

func header() {

	username := "your_username"
	password := "your_password"
	url := "https://api.example.com"

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 添加Basic认证头
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println(resp.Status)

}
