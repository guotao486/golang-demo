/*
 * @Author: GG
 * @Date: 2022-08-18 09:49:12
 * @LastEditTime: 2022-08-18 14:59:29
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\views\index.go
 *
 */
package views

import (
	"errors"
	"golang-demo/blog/golang/service"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

var path string = "/blog/golang"

// index 结构体
// `json:title` 表示返回json格式时对应的key
type IndexData struct {
	Title string `json:title`
	Desc  string `json:desc`
}

// 响应页面
func (*HTMLApi) DemoHtml(w http.ResponseWriter, r *http.Request) {
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

func (*HTMLApi) IndexHtml(w http.ResponseWriter, r *http.Request) {
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

	// 获取参数
	if err := r.ParseForm(); err != nil {
		log.Panicln("表单参数获取失败：", err)
		t.Execute(w, errors.New("系统错误，请联系管理员!!"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	pageSize := 10
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
	}
	// 页面上涉及到的数据必须都有定义
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		t.Execute(w, errors.New("系统错误，请联系管理员!!"))
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
