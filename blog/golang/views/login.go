/*
 * @Author: GG
 * @Date: 2022-08-18 15:25:13
 * @LastEditTime: 2022-08-18 16:48:58
 * @LastEditors: GG
 * @Description:login
 * @FilePath: \golang-demo\blog\golang\views\login.go
 *
 */
package views

import (
	"golang-demo/blog/golang/config"
	"html/template"
	"log"
	"net/http"
	"os"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	t := template.New("login.html")
	// 当前路径
	viewPath, _ := os.Getwd()
	login := viewPath + "/blog/golang/template/login.html"
	footer := viewPath + "/blog/golang/template/layout/footer.html"
	header := viewPath + "/blog/golang/template/layout/header.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(login, footer, header)
	if err != nil {
		log.Panicln(err)
	}

	t.Execute(w, config.Cfg.Viewer)

}
