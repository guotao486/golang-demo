/*
 * @Author: GG
 * @Date: 2022-08-17 14:49:25
 * @LastEditTime: 2022-08-17 15:09:33
 * @LastEditors: GG
 * @Description:README
 * @FilePath: \golang-demo\gorm\index.go
 *
 */
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "go_db:123456@tcp(110.40.208.203:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
