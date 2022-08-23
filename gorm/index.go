/*
 * @Author: GG
 * @Date: 2022-08-17 14:49:25
 * @LastEditTime: 2022-08-23 16:08:42
 * @LastEditors: GG
 * @Description:README
 * @FilePath: \golang-demo\gorm\index.go
 *
 */
package main

import (
	"golang-demo/gorm/models"
	mysqlDB "golang-demo/gorm/mysql"
	"golang-demo/gorm/service"
)

func main() {
	userAtion()
	// creditCardAtion()
}

func userAtion() {
	db := mysqlDB.DB
	// service.UserCreateTable(db) // 创建表
	// service.UserDropTable(db) // 删除表
	// service.UserCreate(db) // 插入数据
	// service.UserSelectCreate(db) // 插入数据，只更新指定字段
	// service.UserOmitCreate(db) // 插入数据，过滤指定字段
	// service.UserBatchCreate(db) // 批量插入数据
	// service.UserMapCreate(db) // 根据map内容插入数据
	// service.UserCreditCardCreate(db) // 关联数据插入
	service.UserDefaultCreate(db) // 默认值
}

func creditCardAtion() {
	mysqlDB.DB.AutoMigrate(&models.CreditCard{})
}
