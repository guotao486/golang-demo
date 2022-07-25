/*
 * @Author: GG
 * @Date: 2022-07-25 17:34:42
 * @LastEditTime: 2022-07-25 17:46:10
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_init.go
 *
 */
package main

import "fmt"

func init() {
	fmt.Println("init")
}
func init() {
	fmt.Println("init2")
}
func main1() {
	fmt.Println("main...")
}
