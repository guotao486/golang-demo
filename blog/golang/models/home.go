/*
 * @Author: GG
 * @Date: 2022-08-17 17:24:59
 * @LastEditTime: 2022-08-17 17:25:22
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\models\home.go
 *
 */
package models

import "golang-demo/blog/golang/config"

type HomeData struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
