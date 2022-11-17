/*
 * @Author: GG
 * @Date: 2022-11-17 10:08:28
 * @LastEditTime: 2022-11-17 14:51:08
 * @LastEditors: GG
 * @Description: 阶梯利润
 * @FilePath: \练手算法\pro2\main.go
 *
 */
package main

import "fmt"

// 企业发放的奖金根据利润提成。利润(I)低于或等于 10 万元时，奖金可提成 10%；利润高于 10 万元，低于 20 万元，低于 10 万元的部分按 10% 提成，高于 10 万元的部分，可提成 7.5%。

// 20 万到 40 万之间时，高于 20 万元的部分，可提成 5%；40 万到 60 万之间时高于 40 万元的部分，可提成 3%；60 万到 100 万之间时，高于 60 万元的部分，可提成 1.5%，高于 100 万元时，超过 100 万元的部分按 1% 提成。

// 从键盘输入当月利润 I，求应发放奖金总数？
func main() {
	var I float32 = 0.0
	fmt.Print("输入利润：")
	fmt.Scanf("%f\n", &I)
	test1(I)
	test2(I)
}

func test1(I float32) {

	var bonus float32 = 0.0

	// fallthrough 关键字，取消switch的break，执行下一个case
	switch {
	case I > 1000000:
		bonus = (I - 1000000) * 0.01
		I = 1000000
		fallthrough
	case I > 600000:
		bonus += (I - 600000) * 0.015
		I = 600000
		fallthrough
	case I > 400000:
		bonus += (I - 400000) * 0.03
		I = 400000
		fallthrough
	case I > 200000:
		bonus += (I - 200000) * 0.05
		I = 200000
		fallthrough
	case I > 100000:
		bonus += (I - 100000) * 0.075
		I = 100000
		fallthrough
	default:
		bonus += I * 0.1
	}
	fmt.Println("TEST 1 提成总计：", bonus)
}
func test2(I float32) {
	var bonus float32 = 0.0
	config := make(map[float32]float32)
	config[1000000] = 0.01
	config[600000] = 0.015
	config[400000] = 0.03
	config[200000] = 0.05
	config[100000] = 0.075
	config[0] = 0.1

	for k, v := range config {
		if I > k {
			bonus += (I - k) * v
			I = k
		}
	}
	fmt.Println("TEST 2 提成总计：", bonus)
}
