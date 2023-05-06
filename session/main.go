package main

import "src/demo/session/session"

// 全局session管理器
var globalSessions *session.Manager

// 全局初始化session管理器
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	// 启动GC回收过期资源
	go globalSessions.GC()
}

func main() {

}
