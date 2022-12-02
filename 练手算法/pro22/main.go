/*
 * @Author: GG
 * @Date: 2022-12-02 16:34:48
 * @LastEditTime: 2022-12-02 16:52:10
 * @LastEditors: GG
 * @Description: 分数序列之和
 * @FilePath: \练手算法\pro22\main.go
 *
 */
package main

import "fmt"

//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13…求出这个数列的前 20 项之和。
//请抓住分子与分母的变化规律。分母的变化规律为分母的值为前两项分母的和，分子的值为同样也是前两项分子的和。
func main() {
	var num = 20
	var a, b, s = 2.0, 1.0, 0.0

	for i := 0; i < num; i++ {
		s += a / b
		a, b = a+b, a
	}
	fmt.Printf("s: %v\n", s)
}
