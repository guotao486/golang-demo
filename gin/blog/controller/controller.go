/*
 * @Author: GG
 * @Date: 2022-09-03 11:44:33
 * @LastEditTime: 2022-09-05 16:20:16
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
	"net/http"

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
