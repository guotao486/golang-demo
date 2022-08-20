/*
 * @Author: GG
 * @Date: 2022-08-17 17:22:46
 * @LastEditTime: 2022-08-19 11:53:09
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\models\category.go
 *
 */
package models

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
