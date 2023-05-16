### 任务调度器

#### 初始化和注册任务
```
/*
 * @Author: GG
 * @Date: 2023-05-12 09:51:08
 * @LastEditTime: 2023-05-16 14:43:42
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\main.go
 *
 */
package main

import (
	"context"
	"net/http"
	"task/controller"
	"task/entity"
	"task/task"
	"task/task/provider/mysql"
)

var GlobalTask *task.Manager

func init() {
	config := mysql.MysqlConfig{
		TableName: "cron_task",
		Username:  "root",
		Password:  "123456",
		Host:      "127.0.0.1",
		Port:      3306,
		Database:  "go_demo",
	}
	mysql.InitDB(config)
	taskM, err := task.NewManager("mysql", 5)
	if err != nil {
		panic(err)
	}

    // 开始任务调度
	go taskM.Run(context.Background())

    // 注册任务主体
	task.EntityRegister("orderExpire", entity.NewOrderExpire())
	task.EntityRegister("orderCreate", entity.NewOrderCreate())
	GlobalTask = taskM
}

func main() {
	index := controller.NewIndexController(GlobalTask)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,world"))
	})
	http.HandleFunc("/create", index.CreateOrder) // 创建订单，发送添加积分任务
	http.HandleFunc("/expire", index.OrderExpire) // 订单失效任务

	http.ListenAndServe(":8088", nil)
}

```

#### 注册任务
```
    data := make(map[string]interface{})
	// 定时任务
	index.task.AddIntervalTask("orderCreate", "订单成功，发送积分", "orderCreate", 10*time.Second, data)
	// 延时任务
	index.task.AddDelayTask("orderExpire", "订单失效", "orderExpire", 60*time.Second, data)
	// 立即执行
	index.task.AddSimpleTask("orderCreate", "订单成功通知", "orderCreate", data)
	// 定时执行
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-05-16 16:15:00")
	index.task.AddTimerTask("orderCreate", "订单售后", "orderCreate", t, data)
```