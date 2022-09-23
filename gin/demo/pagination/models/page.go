/*
 * @Author: GG
 * @Date: 2022-09-23 16:09:38
 * @LastEditTime: 2022-09-23 16:11:44
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\models\page.go
 *
 */
package models

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
