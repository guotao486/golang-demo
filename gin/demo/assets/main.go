/*
 * @Author: GG
 * @Date: 2022-08-31 15:59:44
 * @LastEditTime: 2022-08-31 16:05:56
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\demo\assets\main.go
 *
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Static("/assets", "./assets")                        // 加载文件夹下的静态文件
	e.StaticFS("/more_static", http.Dir("./"))             // 展示目录下文件
	e.StaticFile("/readme.ico", "./README.md")             // 直接访问文件，相对路径
	e.StaticFileFS("/readme", "README.md", http.Dir("./")) // 访问指定目录的指定文件， 绝对路径

	e.Run()
}
