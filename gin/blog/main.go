/*
 * @Author: GG
 * @Date: 2022-09-02 17:24:11
 * @LastEditTime: 2022-09-03 10:16:43
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\blog\main.go
 *
 */
package main

import (
	"golang-demo/gin/blog/dao"
	"golang-demo/gin/blog/models"
)

func main() {
	user := models.User{
		Username: "tom",
		Password: "password",
	}
	dao.Mgr.AddUser(&user)
}
