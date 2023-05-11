## 自定义session管理器

### 内存存储
```
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
```
### mysql
```
package main

import (
	"golang-demo/session/controller"
	"golang-demo/session/session"
	"golang-demo/session/session/mysql"
	"net/http"
)

// 全局session管理器
var GlobalSessions *session.Manager

// 全局初始化session管理器
func init() {
	config := mysql.MysqlConfig{
		TableName: "sessions",
		Username:  "root",
		Password:  "123456",
		Host:      "127.0.0.1",
		Port:      3306,
		Database:  "go_demo",
	}
	mysql.InitDB(config)
	GlobalSessions, _ = session.NewManager("mysql", "gosessionid", 3600, 0)
	// 启动GC回收过期资源
	go GlobalSessions.GC()
}

func main() {
	index := controller.NewIndexController(GlobalSessions)
	http.HandleFunc("/login", index.Login)
	http.HandleFunc("/home", index.Home)

	http.ListenAndServe(":8080", nil)
}
```
### controller
```
package controller

import (
	"fmt"
	"golang-demo/session/session"
	"net/http"
)

type IndexController struct {
	session *session.Manager
}

func NewIndexController(sessionManager *session.Manager) *IndexController {
	return &IndexController{
		session: sessionManager,
	}
}

func (index *IndexController) Login(w http.ResponseWriter, r *http.Request) {
	sess := index.session.SessionStart(w, r)
	sess.Set("username", "tom")
}

func (index *IndexController) Home(w http.ResponseWriter, r *http.Request) {
	sess := index.session.SessionStart(w, r)
	username := sess.Get("username")
	fmt.Printf("username: %v\n", username)
}
```