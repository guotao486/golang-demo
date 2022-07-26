/*
 * @Author: GG
 * @Date: 2022-07-26 21:04:09
 * @LastEditTime: 2022-07-26 21:07:25
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goroutines_select.go
 *
 */
package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main1() {
	go func() {
		chanInt <- 100
		chanStr <- "php"

		// 不关闭的话 select 会一直读取，返回默认值
		close(chanInt)
		close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("r: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("r: %v\n", r)
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second)
	}
}
