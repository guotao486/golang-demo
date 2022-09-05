/*
 * @Author: GG
 * @Date: 2022-09-02 17:26:55
 * @LastEditTime: 2022-09-05 17:23:10
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
	Login(username string) models.User

	// 博客操作
	AddPost(post *models.Post)
	GetAllPost() []models.Post
	GetPost(pid int) models.Post
}

type manager struct {
	db *gorm.DB
}
type demo struct {
	n string
}

var Mgr Manager

func init() {
	dsn := "go_gin_blog_demo:123456@tcp(110.40.208.203:3306)/go_gin_blog_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	// db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}

func (mgr *manager) AddUser(user *models.User) {
	mgr.db.Create(&user)
	fmt.Printf("user: %v\n", user)
}
func (mgr *manager) Login(username string) (user models.User) {
	mgr.db.Where("username=?", username).First(&user)
	return user
}
func (mgr *manager) AddPost(post *models.Post) {
	mgr.db.Create(&post)
}
func (mgr *manager) GetAllPost() []models.Post {
	var posts = make([]models.Post, 10)
	mgr.db.Find(posts)
	return posts
}
func (mgr *manager) GetPost(pid int) (post models.Post) {
	mgr.db.First(&post, pid)
	return
}
