/*
 * @Author: GG
 * @Date: 2022-09-03 11:44:33
 * @LastEditTime: 2022-09-03 11:46:27
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\blog\controller\controller.go
 *
 */
package controller

import (
	"golang-demo/gin/blog/dao"
	"golang-demo/gin/blog/models"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "user.html", nil)
}
