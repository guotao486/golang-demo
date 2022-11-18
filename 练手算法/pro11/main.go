/*
 * @Author: GG
 * @Date: 2022-11-18 16:41:21
 * @LastEditTime: 2022-11-18 16:53:53
 * @LastEditors: GG
 * @Description: 求水仙花数
 * @FilePath: \练手算法\pro11\main.go
 *
 */
package main

import "fmt"

// 打印出所有的 “水仙花数”，所谓 “水仙花数” 是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个 “水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方。
// i*i*i + j*j*j + k*k*k = num
func main() {
	for num := 100; num < 1000; num++ {
		i := num / 100
		j := num / 10 % 10
		k := num % 10
		if i*i*i+j*j*j+k*k*k == num {
			fmt.Printf("%d^3 + %d^3 + %d^3 = %d\n", i, j, k, num)
		}
	}
}
