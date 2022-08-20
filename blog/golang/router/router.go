/*
 * @Author: GG
 * @Date: 2022-08-18 09:46:59
 * @LastEditTime: 2022-08-20 10:59:42
 * @LastEditors: GG
 * @Description: router 路由文件
 * @FilePath: \golang-demo\blog\golang\router\router.go
 *
 */
package router

import (
	"golang-demo/blog/golang/api"
	"golang-demo/blog/golang/views"
	"net/http"
)

// 加载静态资源
func resource() {
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("blog/golang/public/resource/"))))
}

func Router() {
	// 静态资源
	resource()
	// 路由， 响应路径
	// api 返回json
	// http.HandleFunc("/", api.JSONAPI.Index)
	// 页面响应
	http.HandleFunc("/", views.HTML.IndexHtml)
	http.HandleFunc("/demo.html", views.HTML.DemoHtml)
	http.HandleFunc("/index.html", views.HTML.IndexHtml)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/writing", views.HTML.Writing)

	// api
	http.HandleFunc("/api/v1/login", api.JSONAPI.Login)
	http.HandleFunc("/api/v1/post", api.JSONAPI.SaveAndUpdatePost) //修改会保存
	http.HandleFunc("/api/v1/post/", api.JSONAPI.GetPost)          // 获取详情

}
