/*
 * @Author: GG
 * @Date: 2022-08-22 17:18:04
 * @LastEditTime: 2022-08-22 17:53:36
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
