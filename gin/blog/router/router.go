/*
 * @Author: GG
 * @Date: 2022-09-03 11:47:35
 * @LastEditTime: 2022-09-05 16:18:19
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
	e.GET("/", controller.Index)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.RegisterUser)

	e.Run()
}
