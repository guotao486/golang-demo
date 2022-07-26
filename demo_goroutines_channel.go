/*
 * @Author: GG
 * @Date: 2022-07-26 16:50:34
 * @LastEditTime: 2022-07-26 17:57:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goroutines_channel.go
 *
 */
package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func c_show() {
	c := 10
	fmt.Printf("c: %v\n", c)
	channel <- c
	fmt.Println("chan <- channel")
}
func main1() {
	fmt.Println("start")
	go c_show()
	fmt.Println("go start")
	// 不接收的话会阻塞子协程
	v := <-channel
	fmt.Println("<-channel")
	fmt.Printf("v: %v\n", v)
	time.Sleep(time.Second * 3)
}
