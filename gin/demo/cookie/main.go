/*
 * @Author: GG
 * @Date: 2022-09-01 11:05:12
 * @LastEditTime: 2022-09-01 11:16:28
 * @LastEditors: GG
 * @Description: gin cookie
 * @FilePath: \golang-demo\gin\demo\cookie\main.go
 *
 */
package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("/cookie", func(ctx *gin.Context) {
		user := make(map[string]string)
		user["name"] = "tom"
		user["age"] = "18"

		userStr, _ := json.Marshal(user)
		fmt.Printf("userStr: %v\n", string(userStr))
		ctx.SetCookie("userinfo", string(userStr), 60*60, "/", "localhost", false, true)
		ctx.String(200, string(userStr))
	})
	e.GET("/cookie1", func(ctx *gin.Context) {
		userinfo, err := ctx.Cookie("userinfo")
		if err != nil {
			ctx.String(400, err.Error())
		}
		ctx.String(200, userinfo)
	})
	e.Run()
}
