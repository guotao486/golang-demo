/*
 * @Author: GG
 * @Date: 2022-08-30 10:39:16
 * @LastEditTime: 2022-08-30 14:33:27
 * @LastEditors: GG
 * @Description:template
 * @FilePath: \golang-demo\template\main.go
 *
 */
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// 控制台输出
func test1() {
	name := "tom"
	tmlp := "hello,{{.}}"
	t, err := template.New("test").Parse(tmlp)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, name)
	if err != nil {
		log.Fatal(err)
	}
}

// 结构体输出
func test2() {
	type user struct {
		Name string
		Age  string
	}

	myUser := user{
		Name: "tom",
		Age:  "25",
	}
	tmlp := "hello,{{.Name}},Age:{{.Age}}"
	t, err := template.New("test").Parse(tmlp)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, myUser)
	if err != nil {
		log.Fatal(err)
	}
}

// map输出
func test3() {
	user := make(map[string]string)
	user["Name"] = "tom"
	user["age"] = "25"
	tmlp := "hello,{{.Name}},Age:{{.age}}"
	t, err := template.New("test").Parse(tmlp)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}
}

// html
func test4(w http.ResponseWriter, r *http.Request) {
	viewPath, _ := os.Getwd()
	fmt.Printf("viewPath: %v\n", viewPath)
	// d:\php\Go\src\golang-demo

	t, err := template.ParseFiles(viewPath + "/template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	str := []string{"1", "2", "3"}
	t.Execute(w, str)
}
func main() {
	// test1()
	// test2()
	// test3()

	httpServer()
}
func httpServer() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/test4", test4)
	server.ListenAndServe()
}
