/*
 * @Author: GG
 * @Date: 2022-07-26 21:40:47
 * @LastEditTime: 2022-07-26 21:48:44
 * @LastEditors: GG
 * @Description:ticker 根据时间周期执行
 * @FilePath: \golang-demo\demo_ticker.go
 *
 */
package main

import (
	"fmt"
	"time"
)

func main1() {
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {

		if counter >= 5 {
			break
		}
		fmt.Println("ticker 1")
		counter++
	}
	ticker.Stop() //停止

	//协程方式
	ticker2 := time.NewTicker(time.Second)
	chanInt := make(chan int)
	go func() {
		for _ = range ticker2.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("v: %v\n", v)
		sum += v
		if sum >= 10 {
			fmt.Printf("sum: %v\n", sum)
			break
		}

	}
	ticker2.Stop()
}
