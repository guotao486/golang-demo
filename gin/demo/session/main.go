/*
 * @Author: GG
 * @Date: 2022-09-01 11:28:04
 * @LastEditTime: 2022-09-01 16:07:49
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\demo\session\main.go
 *
 */
package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	// 基于cookie的存储引擎, 可更换其他的
	// redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store := cookie.NewStore([]byte("secret"))

	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	e.Use(sessions.Sessions("mySession", store))

	e.GET("/set", func(c *gin.Context) {
		user := make(map[string]string)
		user["name"] = "tom"
		user["age"] = "25"
		userStr, _ := json.Marshal(user)
		//session 使用的编解码器是自带的gob，所以存储类似： struct、map 这些对象时需要先注册对象
		gob.Register(user)
		// 初始化session对象
		session := sessions.Default(c)
		session.Options(sessions.Options{MaxAge: 10})
		// 设置session值

		session.Set("userinfo", user)
		session.Set("userinfo2", userStr)
		sessionStr := session.Get("userinfo")
		// 保存session值
		session.Save()
		c.String(200, "session:%s", sessionStr)
		// c.Request.URL.Path = "/get"
		// e.HandleContext(c)
	})

	e.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionStr := session.Get("userinfo")
		sessionStr2 := session.Get("userinfo2") // json 取出是一个 []uint8
		fmt.Printf("sessionStr2: %T\n", sessionStr2)
		u := sessionStr2.([]uint8) // 类型断言转换一下

		fmt.Printf("u: %v\n", u)
		c.JSON(200, gin.H{
			"userinfo":  sessionStr,
			"userinfo2": string(u), // 转成string
		})
	})

	e.GET("/delete", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("userinfo")
		session.Save()
		sessionStr := session.Get("userinfo")
		c.JSON(200, gin.H{
			"userinfo delete": sessionStr,
		})
	})

	e.GET("/clear", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("userinfo", "userinfo")
		session.Set("name", "abcd")
		// session.Clear()
		session.Save()
		sessionStr := session.Get("userinfo")
		sessionStr2 := session.Get("name")
		c.JSON(200, gin.H{
			"userinfo clear": sessionStr,
			"name clear":     sessionStr2,
		})
	})
	e.Run()
}
