/*
 * @Author: GG
 * @Date: 2022-08-19 14:10:15
 * @LastEditTime: 2022-08-20 10:36:08
 * @LastEditors: GG
 * @Description:post detail
 * @FilePath: \golang-demo\blog\golang\views\post.go
 *
 */
package views

import (
	"errors"
	"golang-demo/blog/golang/service"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	t := template.New("detail.html")
	viewPath, _ := os.Getwd()
	detail := viewPath + "/blog/golang/template/detail.html"
	header := viewPath + "/blog/golang/template/layout/header.html"
	footer := viewPath + "/blog/golang/template/layout/footer.html"

	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t.ParseFiles(detail, header, footer)

	s := r.URL.Path

	s = strings.TrimPrefix(s, "/p/")

	// 判断是否.html 结尾
	if !strings.HasSuffix(s, ".html") {
		return
	}
	pIdStr := strings.TrimSuffix(s, ".html")

	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println(err)
		t.Execute(w, errors.New("不识别此请求路径"))
		return
	}

	detailResponse, err := service.GetPostDetail(pId)

	if err != nil {
		log.Println(err)
		t.Execute(w, err)
		return
	}

	t.Execute(w, detailResponse)
}

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	t := template.New("writing.html")
	viewPath, _ := os.Getwd()
	writing := viewPath + "/blog/golang/template/writing.html"
	t.ParseFiles(writing)
	writingRes, err := service.Writing()
	if err != nil {
		log.Println(err)
		t.Execute(w, err)
		return
	}
	err = t.Execute(w, writingRes)
	log.Println(err)
}
