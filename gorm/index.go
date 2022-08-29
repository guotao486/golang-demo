/*
 * @Author: GG
 * @Date: 2022-08-17 14:49:25
 * @LastEditTime: 2022-08-29 09:31:27
 * @LastEditors: GG
 * @Description:README
 * @FilePath: \golang-demo\gorm\index.go
 *
 */
package main

import (
	"golang-demo/gorm/models"
	mysqlDB "golang-demo/gorm/mysql"
)

func main() {
	// dropTable()
	// createTable()
	userAtion()
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
	// service.UserDefaultCreate(db) // 默认值

	// service.UserFind(db) // 检索单条数据
	// service.UserQueryFirstById(db) // 根据主键
	// service.UserQueryFind(db) // 查询所有数据
	// service.UserQueryWhereFind(db) // 根据where查询

	// service.UserSave(db) // 更新数据
	// service.UserUpdate(db) // 更新单列
	// service.UserUpdates(db)
	// service.UserDelete(db) // 删除，自动判断是否软删除
	// service.UserUnscoped(db) // 查询软删除数据
	// service.UserDrop(db) // 永久删除
	// service.UserRaw(db) // 原生sql

	// service.UserRelatedCreate(db) // 关联model 插入数据,全部是新数据
	// service.UserRelatedCreate2(db) // 给company 新增员工
	// service.UserRelatedCreate3(db) // 给用户新增一条信用卡信息
	// service.UserRelatedUpdate(db) // 修改
	// service.UserRelateSkipUpsert(db)
	// service.UserRelateAssociation(db) // 查询关联数据
	// service.UserRelateAssociationAppend(db) // 添加关联数据
	// service.UserRelateAssociationReplace(db) // 替换关联数据
	// service.UserRelateAssociationDelete(db) // 删除关联
	// service.UserRelateAssociationClear(db) // 清空关联关系
	// service.UserRelateAssociationDrop(db) // 删除源数据并删除关联数据
	// service.UserRelateJoins(db) // 预加载 一对一
	// service.UserRelatePreload(db) // 预加载
}
func createTable() {
	mysqlDB.DB.AutoMigrate(&models.User{}, &models.CreditCard{}, &models.Company{}, &models.Language{})
}
func dropTable() {
	mysqlDB.DB.Migrator().DropTable(&models.User{}, &models.CreditCard{}, &models.Company{})
}
