/*
 * @Author: GG
 * @Date: 2022-08-19 10:39:23
 * @LastEditTime: 2022-08-19 11:56:34
 * @LastEditors: GG
 * @Description:category
 * @FilePath: \golang-demo\blog\golang\views\category.go
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
	"strings"
	"text/template"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	t := template.New("category.html")
	// 当前路径
	viewPath, _ := os.Getwd()

	category := viewPath + path + "/template/category.html"
	header := viewPath + path + "/template/layout/header.html"
	footer := viewPath + path + "/template/layout/footer.html"
	main := viewPath + path + "/template/home.html"
	personal := viewPath + path + "/template/layout/personal.html"
	postlist := viewPath + path + "/template/layout/post-list.html"
	pagination := viewPath + path + "/template/layout/pagination.html"

	// 模板需要使用的函数
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(category, header, footer, main, personal, postlist, pagination)
	if err != nil {
		log.Println(err)
	}

	// 获取url 路径
	//http://localhost:8080/c/1  1参数 分类的id
	s := r.URL.Path // `/c/1`
	cIdStr := strings.TrimPrefix(s, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		log.Println("不识别此请求路径")
		t.Execute(w, errors.New("不识别此请求路径!"))
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("表单数据获取错误")
		t.Execute(w, errors.New("表单数据获取错误!"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	pageSize := 10
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	categoryResult, err := service.GetPostByCategoryId(cId, page, pageSize)
	if err != nil {
		t.Execute(w, err)
		return
	}

	t.Execute(w, categoryResult)
}
