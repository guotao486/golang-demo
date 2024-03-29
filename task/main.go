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

	go taskM.Run(context.Background())

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
