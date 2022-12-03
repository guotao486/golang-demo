/*
 * @Author: GG
 * @Date: 2022-12-03 11:06:59
 * @LastEditTime: 2022-12-03 11:57:21
 * @LastEditors: GG
 * @Description: 洗牌算法
 * @FilePath: \练手算法\pro29\main.go
 *
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 洗牌算法，即将原来的顺序打乱，组成新的随机排序的顺序
func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		shuffle(intArr)
		fmt.Printf("intArr: %v\n", intArr)
	}
}

func shuffle(arr []int) {

	var i, j int
	var temp int
	for i = len(arr) - 1; i > 0; i-- {
		rand.Seed(time.Now().UnixNano()) // 随机基值，同一个基值程序中断后 随机数会一样
		j = rand.Intn(i + 1)             // i+1为随机最大数
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}
