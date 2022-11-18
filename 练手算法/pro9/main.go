/*
 * @Author: GG
 * @Date: 2022-11-18 15:37:40
 * @LastEditTime: 2022-11-18 15:57:32
 * @LastEditors: GG
 * @Description: 兔子繁衍算法
 * @FilePath: \练手算法\pro9\main.go
 *
 */
package main

import "fmt"

// 古典问题：有一对兔子，从出生后第 3 个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
func main() {
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
	var m int = 12
	var x, y int = 0, 1

	for i := 0; i < m; i++ {
		x, y = y, x+y
		fmt.Printf("第%d个月，有%d对兔子", (i + 1), x)
		fmt.Printf("\n")
	}
}
