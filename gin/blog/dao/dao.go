/*
 * @Author: GG
 * @Date: 2022-09-02 17:26:55
 * @LastEditTime: 2022-09-02 17:27:58
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\blog\dao\dao.go
 *
 */
package dao

import (
	"fmt"
	"golang-demo/gin/blog/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "go_gin_blog_demo:123456@tcp(110.40.208.203:3306)/go_gin_blog_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
}

func (mgr *manager) AddUser(user *models.User) {
	mgr.db.Create(&user)
	fmt.Printf("user: %v\n", user)
}
