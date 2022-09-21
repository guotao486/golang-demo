<!--
 * @Author: GG
 * @Date: 2022-09-21 14:29:43
 * @LastEditTime: 2022-09-21 16:48:17
 * @LastEditors: GG
 * @Description: 
 * @FilePath: \logrus\REMADE.MD
 * 
-->
## logrus log日志库

 - GitHub访问地址：https://github.com/sirupsen/logrus
 - 完全兼容golang标准库日志模块：logrus拥有六种日志级别：debug、info、warn、error、fatal和panic，这是golang标准库日志模块的API的超集。如果您的项目使用标准库日志模块，完全可以以最低的代价迁移到logrus上。
 - 可扩展的Hook机制：允许使用者通过hook的方式将日志分发到任意地方，如本地文件系统、标准输出、logstash、elasticsearch或者mq等，或者通过hook定义日志内容和格式等。
 - 可选的日志输出格式：logrus内置了两种日志格式，JSONFormatter和TextFormatter，如果这两个格式不满足需求，可以自己动手实现接口Formatter，来定义自己的日志格式。
 - Field机制：logrus鼓励通过Field机制进行精细化的、结构化的日志记录，而不是通过冗长的消息来记录日志。


#### 安装logrus
```
go get github.com/sirupsen/logrus
```

##### 实例
```
packege main

import log "github.com/sirupsen/logrus"

func main(){
    log.Info("logrus")
}
```
##### 结果
```
time="2022-09-15T11:32:40+08:00" level=info msg="logrus"

```
---

## logrus日志级别
> logrus 支持以下级别 `Info`、`Debug`、`Warn`、`Error`、`Panic`、`Fatal`

```
package main

import log "github.com/sirupsen/logrus"

func main(){
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
}
```
#### 运行结果
```
time="2022-09-21T14:55:32+08:00" level=error msg=error
time="2022-09-21T14:55:33+08:00" level=panic msg=panic
panic: (*logrus.Entry) 0xc00009a000

goroutine 1 [running]:
github.com/sirupsen/logrus.(*Entry).log(0xc00010e000, 0x0, {0xc00008a17a, 0x5})
	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/entry.go:260 +0x47e
github.com/sirupsen/logrus.(*Entry).Log(0xc00010e000, 0x0, {0xc00007df10?, 0x11?, 0xc000037f30?})
	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/entry.go:304 +0x4f
github.com/sirupsen/logrus.(*Logger).Log(0xc00010a000, 0x0, {0xc00007df10, 0x1, 0x1})
	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/logger.go:204 +0x65
github.com/sirupsen/logrus.(*Logger).Panic(...)
	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/logger.go:253
github.com/sirupsen/logrus.Panic(...)
	D:/php/Go/pkg/mod/github.com/sirupsen/logrus@v1.9.0/exported.go:129
main.test2()
	d:/php/Go/src/golang-demo/logrus/main.go:32 +0x165
main.main()
	d:/php/Go/src/golang-demo/logrus/main.go:15 +0x17
exit status 2
```

## 日志格式
> logrus支持两种日志格式，普通文本和json格式，也可以自己定义

```
package main

import log "github.com/sirupsen/logrus"

func main(){
    // 文本格式
    log.SetFormatter(&log.TextFormatter{})
	log.Info("text") 
    //输出 time="2022-09-21T15:15:59+08:00" level=info msg=text

    // json格式
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("json") 
    //输出 {"level":"info","msg":"json","time":"2022-09-21T15:15:59+08:00"}
}
```

## 日志输出
> logrus日志输出有多种方式，可以控制台、文件等。

```
package main

import (
    os

    log "github.com/sirupsen/logrus"
)

func main(){
    // 设置输出到控制台
    log.SetOutput(os.Stdout)
    log.Info("stdout")

    // 设置输出到文件
    file,err:=os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
    if err == nil {
        log.SetOutput(file)
    } else {
        log.Info("日志文件创建失败")
    }
    log.Info("file")
   
}
```

## logrus Fields 结构化记录日志
> logrus不推荐使用冗长的消息来记录运行信息，它推荐使用 `Fields` 来进行精细化的、结构化的信息记录。


```
pacakge main

import log "github.com/sirupsen/logrus"

func main(){
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
```
#### 输出结果
```
time="2022-09-21T16:09:54+08:00" level=info msg="事件发送失败" event="下订单" key=1001 topic=order

time="2022-09-21T16:09:54+08:00" level=info msg="事件发送失败" code=400 event="下订单" key=1001 topic=order uid=10

time="2022-09-21T16:09:54+08:00" level=info msg="key info " key=1001

```

## 输出文件路径、函数名称和行号
> 使用log.SetReportCaller(true)，进行设置，默认是false不输出

```
package main

import log "github.com/sirupsen/logrus"

func main(){
    log.SetReportCaller(true)
    log.Info("info")
}
```
#### 输出
```
time="2022-09-21T16:34:56+08:00" level=info msg=info func=main.test6 file="d:/php/Go/src/golang-demo/logrus/main.go:119"
```

## Logger
> logger是一种相对高级的用法, 对于一个大型项目, 往往需要一个全局的 `logrus` 实例，即logger对象来记录项目所有的日志。

```
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

// 输出 {"event":"下订单","key":1001,"level":"fatal","msg":"事件发送失败","time":"2022-09-21T16:46:28+08:00","topic":"order"}
exit status 1
```

## 分割本地日志文件
> `logrus`本身不带日志本地文件分割功能，但是我们可以通过`file-rotatelogs`进行日志本地文件分割。每次当我们写入日志的时候，`logrus`都会调用`file-rotatelogs`来判断日志是否要进行切分。
```
package main

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func init() {

	path := "message.log"

	/* 日志轮转相关函数

	   `WithLinkName` 为最新的日志建立软连接

	   `WithRotationTime` 设置日志分割的时间，隔多久分割一次

	    WithMaxAge 和 WithRotationCount二者只能设置一个

	    `WithMaxAge` 设置文件清理前的最长保存时间

	    `WithRotationCount` 设置文件清理前最多保存的个数

	*/

	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。

	writer, _ := rotatelogs.New(

		path+".%Y%m%d%H%M",

		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	log.SetOutput(writer)

	//log.SetFormatter(&log.JSONFormatter{})

}

func main() {
	for {
		log.Info("hello, world!")
		time.Sleep(time.Duration(2) * time.Second)
	}
}

```