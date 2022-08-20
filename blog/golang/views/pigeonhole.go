/*
 * @Author: GG
 * @Date: 2022-08-20 13:48:15
 * @LastEditTime: 2022-08-20 14:09:58
 * @LastEditors: GG
 * @Description:pigeonhole views
 * @FilePath: \golang-demo\blog\golang\views\pigeonhole.go
 *
 */
package views

import (
	"golang-demo/blog/golang/service"
	"log"
	"net/http"
	"os"
	"text/template"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {

	t := template.New("pigeonhole.html")
	// 当前路径
	viewPath, _ := os.Getwd()

	index := viewPath + path + "/template/pigeonhole.html"
	header := viewPath + path + "/template/layout/header.html"
	footer := viewPath + path + "/template/layout/footer.html"
	personal := viewPath + path + "/template/layout/personal.html"

	// 模板需要使用的函数
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(index, header, footer, personal)
	if err != nil {
		log.Println(err)
	}

	pigeonholeRes, err := service.FindPostPigeonhole()
	t.Execute(w, pigeonholeRes)

}
