/*
 * @Author: GG
 * @Date: 2022-08-15 16:52:11
 * @LastEditTime: 2022-08-16 11:32:43
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\database\mysql\index.go
 *
 */
package main

import (
	"fmt"
	"golang-demo/database/mysql/curd"
	initDB "golang-demo/database/mysql/init"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := initDB.InitDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Println("初始化成功")
	}
	// 插入
	fmt.Println("插入数据")
	curd.InsertData(initDB.DB)

	// 查询单条
	fmt.Println("查询单条数据：")
	curd.QueryRowDemo(initDB.DB)
	// 查询多条
	fmt.Println("查询多条数据：")
	curd.QueryMuiltDemo(initDB.DB)

}
