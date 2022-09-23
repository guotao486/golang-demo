/*
 * @Author: GG
 * @Date: 2022-09-23 16:26:24
 * @LastEditTime: 2022-09-23 17:34:22
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\controller\user.go
 *
 */
package controller

import (
	"net/http"
	"pagination/dao"
	"pagination/models"
	"pagination/utils"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequestParams(c)

	var user models.User
	list, err := dao.GetAllUsers(&user, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
