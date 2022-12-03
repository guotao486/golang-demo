/*
 * @Author: GG
 * @Date: 2022-12-03 11:59:33
 * @LastEditTime: 2022-12-03 16:06:51
 * @LastEditors: GG
 * @Description: 抢红包
 * @FilePath: \练手算法\pro30\main.go
 *
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	redPackage(10, 100)
}

// 发红包
// count: 红包数量
// money: 红包金额（单位：分)
func redPackage(count, money int) {
	for i := 0; i < count; i++ {
		m := randomMoney(count-i, money)
		money -= m // 红包剩余金额
		fmt.Printf("%d ", m)
	}
}

// 随机红包
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount, remainMonty int) int {
	if remainCount == 1 {
		return remainMonty
	}

	rand.Seed(time.Now().UnixNano())

	max := remainMonty / remainCount * 2 // 红包最大值
	money := rand.Intn(max)
	return money
}
