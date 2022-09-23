/*
 * @Author: GG
 * @Date: 2022-09-23 16:06:42
 * @LastEditTime: 2022-09-23 16:08:50
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagesize\models\user.go
 *
 */
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *User) TableName() string {
	return "user"
}
