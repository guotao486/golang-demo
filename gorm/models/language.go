/*
 * @Author: GG
 * @Date: 2022-08-25 11:02:24
 * @LastEditTime: 2022-08-25 11:18:57
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gorm\models\language.go
 *
 */
package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name string
}
