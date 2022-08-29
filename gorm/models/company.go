/*
 * @Author: GG
 * @Date: 2022-08-24 16:29:42
 * @LastEditTime: 2022-08-25 10:47:40
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gorm\models\company.go
 *
 */
package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name  string
	Users []User
}
