/*
 * @Author: GG
 * @Date: 2022-08-31 11:07:41
 * @LastEditTime: 2022-08-31 15:49:37
 * @LastEditors: GG
 * @Description:gin params demo
 * @FilePath: \golang-demo\gin\demo\params\main.go
 *
 */

/* default默认值 只检查字段是否存在，不存在才返回默认值，若存在即使是空值也返回空值 */
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/* 参数获取 */
// get /login?name=tom&age=25
func Login(c *gin.Context) {

	name := c.Query("name")
	age := c.DefaultQuery("age", "18")
	c.HTML(200, "login.html", gin.H{
		"name": name,
		"age":  age,
	})
}

// post
func DoLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "password1111")
	hobby := c.PostFormArray("hobby")
	c.HTML(200, "index.html", gin.H{
		"username": username,
		"password": password,
		"hobby":    hobby,
	})
}

// resultApi 风格
// 地址参数   localhost:8080/user/1,获取 1
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})
}

func GetPage(c *gin.Context) {
	p := c.Param("p")
	ps := c.Param("ps")
	c.JSON(200, gin.H{
		"p":  p,
		"ps": ps,
	})
	// c.JSONP(200, gin.H{
	// 	"p":  p,
	// 	"ps": ps,
	// })
}

/* 参数获取end */

/* 参数绑定 */
// UserForm struct
// post
type User struct {
	Username string   `form:"username" binding:"required"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func Register(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(200, "User:%s", user)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

// get
type User2 struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func GetBind(c *gin.Context) {
	var user User2
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "User:%s", user)
}

// uri
// resultApi 地址参数绑定
type Pages struct {
	Page     int `uri:"page"`
	PageSize int `uri:"pagesize" default:"10"`
}

func UriBind(c *gin.Context) {
	var pages Pages
	err := c.ShouldBindUri(&pages)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "Pages:%d", pages)
}

/* 参数绑定end */
func main() {
	e := gin.Default()
	viewPath, _ := os.Getwd()

	e.LoadHTMLGlob(viewPath + "/gin/demo/templates/params/*")
	e.GET("/login", Login)
	e.POST("/login", DoLogin)
	e.GET("/user/:id", GetUser)
	e.GET("/page/:p/:ps", GetPage)

	// 数据绑定
	e.GET("/register", GoRegister)
	e.POST("/register", Register)
	e.GET("/getBind", GetBind)
	e.GET("/user_list/:page/:pagesize", UriBind)
	e.Run()
}
