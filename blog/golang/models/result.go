/*
 * @Author: GG
 * @Date: 2022-08-18 16:16:01
 * @LastEditTime: 2022-08-19 09:38:36
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\models\result.go
 *
 */
package models

type Result struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
}
