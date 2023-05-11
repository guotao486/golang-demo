/*
 * @Author: GG
 * @Date: 2023-05-06 11:29:01
 * @LastEditTime: 2023-05-11 17:24:38
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
	"golang-demo/session/session/mysql"
	"net/http"
)

// 全局session管理器
var GlobalSessions *session.Manager

// 全局初始化session管理器
func init() {

	// GlobalSessions, _ = session.NewManager("memory", "gosessionid", 3600, 60)

	config := mysql.MysqlConfig{
		TableName: "sessions",
		Username:  "root",
		Password:  "123456",
		Host:      "127.0.0.1",
		Port:      3306,
		Database:  "go_demo",
	}
	mysql.InitDB(config)
	GlobalSessions, _ = session.NewManager("mysql", "gosessionid", 120, 0)
	// 启动GC回收过期资源
	go GlobalSessions.GC()
}

func main() {
	index := controller.NewIndexController(GlobalSessions)
	http.HandleFunc("/login", index.Login)
	http.HandleFunc("/home", index.Home)

	http.ListenAndServe(":8080", nil)
}
