/*
 * @Author: GG
 * @Date: 2022-09-02 17:04:29
 * @LastEditTime: 2022-09-05 16:54:23
 * @LastEditors: GG
 * @Description:model
 * @FilePath: \golang-demo\gin\blog\models\model.go
 *
 */
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
