/*
 * @Author: GG
 * @Date: 2022-08-31 16:40:03
 * @LastEditTime: 2022-09-02 15:53:16
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\demo\middleware\main.go
 *
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Demo(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func Middleware1(c *gin.Context) {
	fmt.Println("1111")
	c.String(http.StatusOK, "11")
}

func Middleware2(c *gin.Context) {
	fmt.Println("2")
	c.String(http.StatusOK, "2")
}
func Middleware3(c *gin.Context) {
	fmt.Println("3")
	c.String(http.StatusOK, "3")
}

// gin 自带 basicauth中间件
// 模拟私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@gmail.com", "phone": "123456"},
	"austin": gin.H{"email": "austin@gmail.com", "phone": "666666"},
	"lena":   gin.H{"email": "lena@gmail.com", "phone": "654789"},
}

func middleware_test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("<------- middleware_test")
	}
}

func main() {
	// 写日志文件
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// 同时写入文件和输出控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	e := gin.Default()
	e.GET("/demo", Demo)

	// 中间件执行顺序与代码执行顺序有关
	// 路由组使用 gin.BasicAuth()中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	admin := e.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}), Middleware1).Use(Middleware3)

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	admin.GET("/secrets", Middleware2, func(c *gin.Context) {
		// 获取用户名，由BasicAuth中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Printf("user: %v\n", user)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	e.Use(Middleware1)
	// 多个中间件的调度，next() 下一个 abort 中止返回
	e.GET("/test", func(c *gin.Context) {
		fmt.Println("<-------- test")
		c.Next()
		// c.Abort()
		fmt.Println("<----------- test end")
	}, middleware_test())

	// 恢复中间件，异常可恢复执行避免程序中止
	e.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		if err, ok := err.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error:%s", err))
		}
		// 调用abort中止执行，并写入响应码
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	e.GET("/panic", func(c *gin.Context) {
		panic("foo")
	})

	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	e.Run()
}
