/*
 * @Author: GG
 * @Date: 2022-09-23 17:02:19
 * @LastEditTime: 2022-09-24 10:57:01
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\dao\user.go
 *
 */
package dao

import (
	"pagination/config"
	"pagination/models"
)

func GetAllUsers(user *models.User, pagination *models.Pagination) ([]models.User, error) {
	var users []models.User
	var total int64

	// 总数量
	config.DB.Model(&models.User{}).Count(&total)
	pagination.Total = int(total)
	// 总页数
	pagination.GetTotalPage()
	// 分页查询
	// page 1 limit 20   offset 0
	// page 2 limit 20   offset 20
	// page 3 limit 20   offset 40
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuild := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	result := queryBuild.Model(&models.User{}).Where(user).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil

}
