/*
 * @Author: GG
 * @Date: 2022-07-26 21:08:31
 * @LastEditTime: 2022-07-26 21:40:36
 * @LastEditors: GG
 * @Description:timer 延时
 * @FilePath: \golang-demo\demo_timer.go
 *
 */
package main

import (
	"fmt"
	"time"
)

func main1() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1: %v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2: %v\n", t2)

	// 延时等待，可用time.sleep 实现
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("2s后....")

	time.Sleep(time.Second * 2)
	fmt.Println("2s后....")

	<-time.After(time.Second * 2)
	fmt.Println("2s后....")

	// 定时器
	timer3 := time.NewTimer(time.Second * 5)
	go func() {
		<-timer3.C
		fmt.Println("timer3 2 expired")
	}()

	stop := timer3.Stop() //停止定时器
	if stop {
		fmt.Println("timer3 stopped")
	}

	// 定时器时间重置
	timer4 := time.NewTimer(time.Second * 5)
	timer4.Reset(time.Second * 1)
	<-timer4.C
	fmt.Println("timer4 end")

}
