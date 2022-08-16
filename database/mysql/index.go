/*
 * @Author: GG
 * @Date: 2022-08-15 16:52:11
 * @LastEditTime: 2022-08-16 11:08:40
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\database\mysql\index.go
 *
 */
package main

import (
	"fmt"
	initDB "golang-demo/database/mysql/init"
	query "golang-demo/database/mysql/select"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := initDB.InitDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功")
	}
	query.QueryRowDemo(initDB.DB)
	query.QueryMuiltDemo(initDB.DB)
}
