/*
 * @Author: GG
 * @Date: 2022-09-03 11:47:35
 * @LastEditTime: 2022-09-06 15:42:23
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

	e.GET("/post_index", controller.GoPostIndex)
	e.GET("/post", controller.GoAddPost)
	e.POST("/post", controller.AddPost)
	e.GET("/detail", controller.GetPost)
	e.Run()
}
