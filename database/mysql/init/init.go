/*
 * @Author: GG
 * @Date: 2022-08-15 17:33:17
 * @LastEditTime: 2022-08-15 17:37:35
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\database\mysql\init\init.go
 *
 */
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 获得链接 Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。
func open() {
	db, err := sql.Open("mysql", "root:123456@/go_db")
	if err != nil {
		panic(err)
	}
	print(db)
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)
}

//定义一个全局db对象
var db *sql.DB

// 初始化连接
func initDB() (err error) {
	dsn := "go_db:123456@tcp(110.40.208.203:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功")
	}
}
