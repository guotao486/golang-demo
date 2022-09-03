/*
 * @Author: GG
 * @Date: 2022-09-03 11:47:35
 * @LastEditTime: 2022-09-03 11:50:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\blog\router\router.go
 *
 */
package router

import (
	"golang-demo/gin/blog/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	viewPath, _ := os.Getwd()
	e.LoadHTMLGlob(viewPath + "/gin/blog/templates/*")
	e.Static("/assets", viewPath+"/gin/blog/assets")
	e.GET("/users", controller.ListUser)
	e.POST("/users", controller.AddUser)

	e.Run()
}
