/*
 * @Author: GG
 * @Date: 2022-10-10 15:16:58
 * @LastEditTime: 2022-10-11 16:02:15
 * @LastEditors: GG
 * @Description:
 * @FilePath: \swagger\models\model.go
 *
 */
package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);" json:"title" form:"title" example:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100)" json:"des" form:"des" example:"desc" binding:"required"`
	Status string `gorm:"type:varchar(200)" json:"status" form:"status" example:"Active"`
}

type Response struct {
	Msg  string
	Data interface{}
}
