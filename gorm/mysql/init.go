/*
 * @Author: GG
 * @Date: 2022-08-22 14:48:12
 * @LastEditTime: 2022-08-22 14:57:15
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gorm\mysql\init.go
 *
 */
package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm 使用内部的gormDB对象
var DB *gorm.DB

func init() {
	dsn := "go_demo_gorm:123456@tcp(110.40.208.203:3306)/go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
