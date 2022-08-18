/*
 * @Author: GG
 * @Date: 2022-08-18 17:06:16
 * @LastEditTime: 2022-08-18 17:11:15
 * @LastEditors: GG
 * @Description:LoginRes models
 * @FilePath: \golang-demo\blog\golang\models\login.go
 *
 */
package models

type LoginRes struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
