/*
 * @Author: GG
 * @Date: 2022-08-31 16:40:03
 * @LastEditTime: 2022-09-01 11:04:01
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\demo\middleware\main.go
 *
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Demo(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func Middleware1(c *gin.Context) {
	fmt.Println("1")
	c.String(http.StatusOK, "1")
}

func Middleware2(c *gin.Context) {
	fmt.Println("2")
	c.String(http.StatusOK, "2")
}

// gin 自带 basicauth中间件
// 模拟私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@gmail.com", "phone": "123456"},
	"austin": gin.H{"email": "austin@gmail.com", "phone": "666666"},
	"lena":   gin.H{"email": "lena@gmail.com", "phone": "654789"},
}

func main() {
	e := gin.Default()
	e.GET("/demo", Demo)

	// 路由组使用 gin.BasicAuth()中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	admin := e.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}), Middleware1)

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	admin.GET("/secrets", Middleware1, func(c *gin.Context) {
		// 获取用户名，由BasicAuth中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Printf("user: %v\n", user)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	}, Middleware2)
	e.Run()
}
