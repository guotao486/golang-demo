/*
 * @Author: GG
 * @Date: 2022-07-26 16:50:34
 * @LastEditTime: 2022-07-26 21:03:33
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

	// 通道遍历取值
	c := make(chan int)
	c2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Printf("v: %v\n", v)
	}

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for {
		if data, ok := <-c2; ok {
			fmt.Printf("data: %v\n", data)
			fmt.Printf("ok: %v\n", ok)
		} else {
			break
		}
	}

}
