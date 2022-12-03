/*
 * @Author: GG
 * @Date: 2022-12-02 17:53:05
 * @LastEditTime: 2022-12-03 10:07:02
 * @LastEditors: GG
 * @Description: 将数字逆序
 * @FilePath: \练手算法\pro27\main.go
 *
 */
package main

import (
	"fmt"
	"math"
)

// 给一个不多于 5 位的正整数，要求：一、求它是几位数，二、逆序打印出各位数字。
func main() {
	var x int
	fmt.Printf("请输入一个不多于5位数的整数")
	fmt.Scanf("%d", &x)

	for i := 5; i > 0; i-- {
		r := x / int(math.Pow10(i-1)) // 5000/int(10^4) = 0， 5000/int(10^3) = 5

		if r > 0 {
			fmt.Printf("%d 是一个 %d 位数字.\n", x, i)
			out(i, x) // i=4  x=5000
			fmt.Println("")
			break
		}
	}
}

func out(i, x int) {
	if i > 1 {
		out(i-1, x)
	}
	// 5000%10^4/10^3   5000%10000/1000  5000%1000/100
	fmt.Printf("x: %v\n", x)
	fmt.Println("i:,", i, "----", int(math.Pow10(i)))
	fmt.Println("i:", int(math.Pow10(i-1)))
	r := x % int(math.Pow10(i)) / int(math.Pow10(i-1))
	fmt.Printf("%d", r)
}
