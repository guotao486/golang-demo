/*
 * @Author: GG
 * @Date: 2022-08-18 16:16:01
 * @LastEditTime: 2022-08-18 16:17:24
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\models\result.go
 *
 */
package models

type Result struct {
	Error string `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
}
