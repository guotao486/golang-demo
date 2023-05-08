/*
 * @Author: GG
 * @Date: 2023-05-06 11:29:01
 * @LastEditTime: 2023-05-08 15:32:19
 * @LastEditors: GG
 * @Description:
 * @FilePath: \session\main.go
 *
 */
package main

import (
	"golang-demo/session/controller"
	"golang-demo/session/session"
	_ "golang-demo/session/session/memory"
	"net/http"
)

// 全局session管理器
var GlobalSessions *session.Manager

// 全局初始化session管理器
func init() {
	GlobalSessions, _ = session.NewManager("memory", "gosessionid", 3600, 60)
	// 启动GC回收过期资源
	go GlobalSessions.GC()
}

func main() {
	index := controller.NewIndexController(GlobalSessions)
	http.HandleFunc("/login", index.Login)
	http.HandleFunc("/home", index.Home)

	http.ListenAndServe(":8080", nil)
}
