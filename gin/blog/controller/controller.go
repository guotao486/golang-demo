/*
 * @Author: GG
 * @Date: 2022-09-03 11:44:33
 * @LastEditTime: 2022-09-06 15:41:57
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\blog\controller\controller.go
 *
 */
package controller

import (
	"fmt"
	"golang-demo/gin/blog/dao"
	"golang-demo/gin/blog/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := dao.Mgr.Login(username)

	if user.Username == "" {
		c.HTML(http.StatusOK, "login.html", "用户名不存在")
		fmt.Println("用户名不存在")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误")
			fmt.Println("密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}

	c.Redirect(301, "/")

}
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func GoPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := models.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func GetPost(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	post := dao.Mgr.GetPost(pid)
	c.HTML(200, "detail.html", gin.H{
		"Title":   post.Title,
		"Content": template.HTML(post.Content),
	})
}
