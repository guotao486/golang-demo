/*
 * @Author: GG
 * @Date: 2022-08-17 15:13:06
 * @LastEditTime: 2022-08-17 18:05:32
 * @LastEditors: GG
 * @Description: 原生golang 简单博客
 * @FilePath: \golang-demo\blog\golang\main.go
 *
 */
package main

import (
	"encoding/json"
	"golang-demo/blog/golang/config"
	"golang-demo/blog/golang/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var path string = "/blog/golang"

// index 结构体
// `json:title` 表示返回json格式时对应的key
type IndexData struct {
	Title string `json:title`
	Desc  string `json:desc`
}

// 响应对应路径的函数
// http.ResponseWriter 响应对象，http.Request 请求对象
func index(w http.ResponseWriter, r *http.Request) {
	// 设置响应抬头
	w.Header().Set("Content-type", "application/json")

	var indexData IndexData
	indexData.Title = "bolang blog title"
	indexData.Desc = "bolang blog desc"
	jsonStr, _ := json.Marshal(indexData)
	// 响应输出
	// w.Write([]byte("hello world!"))
	w.Write(jsonStr)
}

// 响应页面
func demoHtml(w http.ResponseWriter, r *http.Request) {
	// 实例化模板对象
	t := template.New("demo.html")
	viewPath, _ := os.Getwd()
	// 要解析的模板文件
	t, _ = t.ParseFiles(viewPath + "/blog/golang/template/index.html")
	var indexData IndexData
	indexData.Title = "bolang blog title"
	indexData.Desc = "bolang blog desc"
	// 开始执行解析
	t.Execute(w, indexData)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	// 当前路径
	viewPath, _ := os.Getwd()

	index := viewPath + path + "/template/index.html"
	header := viewPath + path + "/template/layout/header.html"
	footer := viewPath + path + "/template/layout/footer.html"
	main := viewPath + path + "/template/home.html"
	personal := viewPath + path + "/template/layout/personal.html"
	postlist := viewPath + path + "/template/layout/post-list.html"
	pagination := viewPath + path + "/template/layout/pagination.html"

	// 模板需要使用的函数
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(index, header, footer, main, personal, postlist, pagination)
	if err != nil {
		log.Println(err)
	}

	// 页面上涉及到的数据必须都有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}

	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}

	var hr = &models.HomeData{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, hr)
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func main() {
	// 程序入口，一个程序只有一个入口文件
	// web服务，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// 路由， 响应路径
	http.HandleFunc("/", index)
	http.HandleFunc("/demo.html", demoHtml)
	http.HandleFunc("/index.html", indexHtml)

	// server.ListenAndServe 监听端口并启动服务
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
