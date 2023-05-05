/*
 * @Author: GG
 * @Date: 2023-05-04 14:38:05
 * @LastEditTime: 2023-05-05 14:45:07
 * @LastEditors: GG
 * @Description:
 * @FilePath: \xlsx\main.go
 *
 */
package main

import (
	"xlsx/controller"
	"xlsx/dao"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/shopyy_demo"), &gorm.Config{})
	if err != nil {
		panic("数据库连接错误")
	}

	productDao := dao.NewProductDao(db)
	handler := controller.NewProductHandler(*productDao)
	r := gin.Default()
	r.GET("/export", handler.Export)
	r.GET("/import", handler.Import)
	r.Run(":8080")
}
