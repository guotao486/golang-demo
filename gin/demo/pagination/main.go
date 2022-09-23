/*
 * @Author: GG
 * @Date: 2022-09-23 16:04:36
 * @LastEditTime: 2022-09-23 16:57:56
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\main.go
 *
 */
package main

import (
	"pagination/config"
	"pagination/models"
	"pagination/routers"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// config
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&models.User{})
	// 链接池end
	config.DB = db
	r := routers.SetUpRouter()
	r.Run()
}
