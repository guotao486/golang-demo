/*
 * @Author: GG
 * @Date: 2022-08-22 14:48:12
 * @LastEditTime: 2022-08-22 17:32:32
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gorm\mysql\init.go
 *
 */
package mysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm 使用内部的gormDB对象
var DB *gorm.DB

/* db, err := gorm.Open(mysql.New(mysql.Config{
	DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	DefaultStringSize: 256, // string 类型字段的默认长度
	DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
   }), &gorm.Config{}) */

func init() {
	dsn := "go_demo_gorm:123456@tcp(110.40.208.203:3306)/go_demo_gorm?charset=utf8mb4&parseTime=True&loc=Local"
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
	// 链接池end
	DB = db
}
