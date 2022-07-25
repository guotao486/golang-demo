/*
 * @Author: GG
 * @Date: 2022-07-25 17:12:24
 * @LastEditTime: 2022-07-25 17:34:32
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_defer.go
 *
 */
package main

import "fmt"

func defer_f() {
	fmt.Println("defer end....")
}
func main1() {
	fmt.Println("start....")

	defer defer_f()
	fmt.Println("step1")
	fmt.Println("step2")
	fmt.Println("step3")
	fmt.Println("end..")
}
