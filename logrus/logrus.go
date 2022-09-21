/*
 * @Author: GG
 * @Date: 2022-09-21 16:37:37
 * @LastEditTime: 2022-09-21 16:46:06
 * @LastEditors: GG
 * @Description:
 * @FilePath: \logrus\logrus.go
 *
 */
package main

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// log实例
var Log = logrus.New()

// 初始化配置
func init() {
	// 输出终端
	Log.Out = os.Stdout

	// 输出格式
	Log.Formatter = &logrus.JSONFormatter{}

	// 日志级别
	Log.SetLevel(logrus.InfoLevel)
}

func main() {
	var event = "下订单"
	var topic = "order"
	var key = 1001

	//替代方案
	Log.WithFields(log.Fields{
		"event": event,
		"topic": topic,
		"key":   key,
	}).Fatal("事件发送失败")
}
