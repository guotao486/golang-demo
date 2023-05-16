/*
 * @Author: GG
 * @Date: 2023-05-12 10:59:57
 * @LastEditTime: 2023-05-15 10:07:50
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\main_test.go
 *
 */
package main

import (
	"context"
	"fmt"
	"task/task"
	"task/task/provider/mysql"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	fmt.Println("start")
	goleak.VerifyTestMain(m)
}

func TestPoller(t *testing.T) {
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
	taskM.Poll(context.Background())
}
