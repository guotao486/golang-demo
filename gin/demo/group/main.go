/*
 * @Author: GG
 * @Date: 2022-09-01 16:22:52
 * @LastEditTime: 2022-09-01 16:50:22
 * @LastEditors: GG
 * @Description: gin router group
 * @FilePath: \golang-demo\gin\demo\group\main.go
 *
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func f1(c *gin.Context) {
	fmt.Println("f1")
}
func f2(c *gin.Context) {
	fmt.Println("f2")
}
func f3(c *gin.Context) {
	fmt.Println("f3")
}

func main() {
	e := gin.Default()
	v1 := e.Group("/v1")
	{
		v1.GET("/f1", f1)

		v11 := v1.Group("/f2")
		{
			v11.GET("/f3", f3)

		}
	}

	v2 := e.Group("/v2")
	{
		v2.GET("/f2", f2)
	}
	e.Run()
}
