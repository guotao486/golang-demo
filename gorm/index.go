/*
 * @Author: GG
 * @Date: 2022-08-17 14:49:25
 * @LastEditTime: 2022-08-22 17:50:54
 * @LastEditors: GG
 * @Description:README
 * @FilePath: \golang-demo\gorm\index.go
 *
 */
package main

import (
	mysqlDB "golang-demo/gorm/mysql"
	"golang-demo/gorm/service"
)

func main() {
	db := mysqlDB.DB

	// service.UserCreateTable(db) // 创建表
	// service.UserDropTable(db) // 删除表
	service.UserCreate(db) // 插入数据
}
