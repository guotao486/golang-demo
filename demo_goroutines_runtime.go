/*
 * @Author: GG
 * @Date: 2022-07-26 20:40:00
 * @LastEditTime: 2022-07-26 20:44:16
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goroutines_runtime.go
 *
 */
package main

import (
	"fmt"
	"runtime"
)

func r_show(s string) {
	for i := 0; i < 5; i++ {
		if i >= 3 {
			// 退出当前协程
			runtime.Goexit()
		}
		fmt.Println(s)
	}
}
func main1() {

	go r_show("php")
	for i := 0; i < 2; i++ {
		// 让出当前执行，让下一个协程执行
		runtime.Gosched()
		fmt.Println("golang")
	}

	runtime.GOMAXPROCS(1) //修改cpu使用核心数量
}
