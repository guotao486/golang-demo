/*
 * @Author: GG
 * @Date: 2022-08-17 17:24:59
 * @LastEditTime: 2022-08-18 14:33:54
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\models\home.go
 *
 */
package models

import (
	"golang-demo/blog/golang/config"
	"time"
)

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}

func DateDay(t time.Time) string {

	return t.Format("2006-01-02 15:04:05")
}
