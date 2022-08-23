/*
 * @Author: GG
 * @Date: 2022-08-22 17:18:04
 * @LastEditTime: 2022-08-23 16:08:23
 * @LastEditors: GG
 * @Description: User ation
 * @FilePath: \golang-demo\gorm\service\user.go
 *
 */
package service

import (
	"fmt"
	"golang-demo/gorm/models"
	"time"

	"gorm.io/gorm"
)

func UserCreateTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func UserDropTable(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
}

// 创建数据
func UserCreate(db *gorm.DB) {
	user := models.User{
		Name:     "tom",
		Email:    "asdasd@qweq.com",
		Age:      12,
		Birthday: time.Now(),
	}
	result := db.Create(&user)
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
	fmt.Printf("user: %v\n", user)
}

// 创建数据，但只更新select 指定字段
func UserSelectCreate(db *gorm.DB) {
	user := models.User{
		Name:     "tom",
		Email:    "asdasd@qweq.com",
		Age:      12,
		Birthday: time.Now(),
	}

	db.Select("name", "age").Create(&user)
	fmt.Printf("user: %v\n", user)
}

// 创建数据，过滤掉select 指定字段
func UserOmitCreate(db *gorm.DB) {
	user := models.User{
		Name:     "tom",
		Email:    "dasdsa@dsad.com",
		Age:      25,
		Birthday: time.Now(),
	}
	db.Omit("name", "email", "age", "birthday").Create(&user)
	fmt.Printf("user: %v\n", user)
}

// 批量创建
func UserBatchCreate(db *gorm.DB) {
	users := []models.User{
		{Name: "tom"},
		{Name: "tom1"},
		{Name: "tom2"},
	}

	// 一次性
	// db.Create(&users)
	// 分批次
	db.CreateInBatches(users, 1)
	for _, u := range users {
		fmt.Printf("u.ID: %v\n", u.ID)
	}

}

// 根据map类型创建数据
// 不会触发hook和 时间字段追踪
func UserMapCreate(db *gorm.DB) {
	// 单条
	var user = map[string]interface{}{
		"Name": "tom", "Age": 18,
	}

	db.Model(&models.User{}).Create(user)

	// 批量
	var users = []map[string]interface{}{
		{"Name": "tom", "Age": 18},
		{"Name": "tom2", "age": 20},
	}
	db.Model(&models.User{}).Create(users)
}

// 关联数据创建，如果关联值是非零值，这些关联会被 upsert，且它们的 Hook 方法也会被调用
func UserCreditCardCreate(db *gorm.DB) {
	db.Create(&models.User{
		Name: "123",
		CreditCard: models.CreditCard{
			Number: "12312321313",
		},
	})

	// 关联数据空
	db.Create(&models.User{
		Name:       "321",
		CreditCard: models.CreditCard{},
	})

	//db.Omit("CreditCard").Create(&user)

	// 跳过所有关联
	//db.Omit(clause.Associations).Create(&user)
}

// 测试默认值
func UserDefaultCreate(db *gorm.DB) {

	db.Create(&models.User{})
}
