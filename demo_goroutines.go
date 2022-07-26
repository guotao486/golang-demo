/*
 * @Author: GG
 * @Date: 2022-07-26 16:30:09
 * @LastEditTime: 2022-07-26 16:50:07
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goroutines .go
 *
 */
package main

import (
	"fmt"
	"time"
)

func show(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v, m:%v\n", i, msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main1() {
	go show("java")
	show("golang")

}
