/*
 * @Author: GG
 * @Date: 2022-08-18 10:17:36
 * @LastEditTime: 2022-08-18 10:57:19
 * @LastEditors: GG
 * @Description:mysql init
 * @FilePath: \golang-demo\blog\golang\dao\mysql.go
 *
 */
package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dsn := fmt.Sprintf("go_demo_blog:123456@tcp(110.40.208.203:3306)/go_demo_blog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("连接数据库异常")
		log.Panic(err)
	}

	// 最大空闲连接数， 默认2
	db.SetMaxIdleConns(5)
	// 最大连接数，默认不限制
	db.SetMaxOpenConns(100)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Hour)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Hour * 2)

	// 连接
	err = db.Ping()
	if err != nil {
		log.Panicln("数据库连接异常")
		// 关闭
		_ = db.Close()
		log.Panic(err)
	}

	DB = db
}
