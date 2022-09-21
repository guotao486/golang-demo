/*
 * @Author: GG
 * @Date: 2022-09-21 14:27:22
 * @LastEditTime: 2022-09-21 16:45:54
 * @LastEditors: GG
 * @Description:
 * @FilePath: \logrus\main.go
 *
 */
package main

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main1() {
	test7()
}

func test1() {
	log.Info("log") // time="2022-09-21T14:54:27+08:00" level=info msg=log
}

// 日志级别
func test2() {
	// 设置日志级别
	log.SetLevel(log.ErrorLevel)
	// 下面输出，日志级别由低到高，输出情况由日志级别控制
	// 例如：设置日志级别为 ErrorLevel，则 infor、debug、warn不再输出

	log.Info("info")
	log.Debug("debug")
	log.Warn("warn")
	log.Error("error")
	log.Panic("panic")
	log.Fatal("fatal")

	// 	time="2022-09-21T14:55:32+08:00" level=error msg=error
	// time="2022-09-21T14:55:33+08:00" level=panic msg=panic
	// panic: (*logrus.Entry) 0xc00009a000

	// goroutine 1 [running]:
	// github.com/sirupsen/logrus.(*Entry).log(0xc00010e000, 0x0, {0xc00008a17a, 0x5})
	// 	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/entry.go:260 +0x47e
	// github.com/sirupsen/logrus.(*Entry).Log(0xc00010e000, 0x0, {0xc00007df10?, 0x11?, 0xc000037f30?})
	// 	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/entry.go:304 +0x4f
	// github.com/sirupsen/logrus.(*Logger).Log(0xc00010a000, 0x0, {0xc00007df10, 0x1, 0x1})
	// 	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/logger.go:204 +0x65
	// github.com/sirupsen/logrus.(*Logger).Panic(...)
	// 	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/logger.go:253
	// github.com/sirupsen/logrus.Panic(...)
	// 	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/exported.go:129
	// main.test2()
	// 	d:/php/Go/src/golang-demo/logrus/main.go:32 +0x165
	// main.main()
	// 	d:/php/Go/src/golang-demo/logrus/main.go:15 +0x17
	// exit status 2
}

// 日志输出格式，文本或json，也可以自定义
func test3() {
	log.SetFormatter(&log.TextFormatter{})
	log.Info("text") // time="2022-09-21T15:15:59+08:00" level=info msg=text

	log.SetFormatter(&log.JSONFormatter{})
	log.Info("json") // {"level":"info","msg":"json","time":"2022-09-21T15:15:59+08:00"}
}

// 日志输出
func test4() {
	// 输出到控制台
	log.SetOutput(os.Stdout)
	log.Info("stdout")

	// 输出到文件
	file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("日志文件创建失败,原因：", err)
	}
	log.Info("file")
}

// fields 结构化记录日志
func test5() {
	var event = "下订单"
	var topic = "order"
	var key = 1001

	log.WithFields(log.Fields{
		"event": event,
		"topic": topic,
		"key":   key,
	}).Info("事件发送失败")
	// 输出 time="2022-09-21T16:09:54+08:00" level=info msg="事件发送失败" event="下订单" key=1001 topic=order
	var code = 400
	var uid = 10

	log.WithFields(log.Fields{
		"event": event,
		"topic": topic,
		"key":   key,
		"code":  code,
		"uid":   uid,
	}).Info("事件发送失败")
	// 输出 time="2022-09-21T16:09:54+08:00" level=info msg="事件发送失败" code=400 event="下订单" key=1001 topic=order uid=10

	// 单个字段
	log.WithField("key", key).Info("key info ")
	// 输出 time="2022-09-21T16:09:54+08:00" level=info msg="key info " key=1001
}

//显示文件路径、函数名、行号
func test6() {
	log.SetReportCaller(true)
	log.Info("info")
	//time="2022-09-21T16:34:56+08:00" level=info msg=info func=main.test6 file="d:/php/Go/src/golang-demo/logrus/main.go:119"
}
func New() {
	// log实例
	var Log = logrus.New()

	// 初始化配置
	// 输出终端
	Log.Out = os.Stdout

	// 输出格式
	Log.Formatter = &logrus.JSONFormatter{}

	// 日志级别
	Log.SetLevel(logrus.InfoLevel)
}
