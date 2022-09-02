/*
 * @Author: GG
 * @Date: 2022-09-02 17:04:29
 * @LastEditTime: 2022-09-02 17:28:08
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
