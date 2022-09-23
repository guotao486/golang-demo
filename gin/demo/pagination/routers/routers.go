/*
 * @Author: GG
 * @Date: 2022-09-23 16:25:03
 * @LastEditTime: 2022-09-23 17:20:48
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\routers\routers.go
 *
 */
package routers

import (
	"pagination/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	u1 := r.Group("/admin")
	{
		u1.GET("all", controller.GetAllUsers)
		u1.GET("all2/:limit/:page", controller.GetAllUsers)
	}
	return r
}
