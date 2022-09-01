/*
 * @Author: GG
 * @Date: 2022-08-31 09:58:46
 * @LastEditTime: 2022-08-31 10:50:10
 * @LastEditors: GG
 * @Description: gin login demo
 * @FilePath: \golang-demo\gin\demo\login\main.go
 *
 */
package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.HTML(200, "index.html", gin.H{
		"username": username,
		"password": password,
	})
}

func main() {
	e := gin.Default()
	viewPath, _ := os.Getwd()
	e.LoadHTMLGlob(viewPath + "/gin/demo/templates/login/*")
	e.GET("/login", Login)
	e.POST("/login", DoLogin)
	e.Run()
}
