/*
 * @Author: GG
 * @Date: 2022-08-18 15:38:46
 * @LastEditTime: 2022-08-18 15:42:11
 * @LastEditors: GG
 * @Description:user models
 * @FilePath: \golang-demo\blog\golang\models\user.go
 *
 */
package models

import "time"

type User struct {
	Uid      int       `json:"uid"`
	UserName string    `json:"userName"`
	Passwd   string    `json:"passwd"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
