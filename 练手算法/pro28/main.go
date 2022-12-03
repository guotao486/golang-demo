/*
 * @Author: GG
 * @Date: 2022-12-03 10:18:03
 * @LastEditTime: 2022-12-03 10:55:08
 * @LastEditors: GG
 * @Description: 回文数
 * @FilePath: \练手算法\pro28\main.go
 *
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

// 一个 5 位数，判断它是不是回文数。即 12321 是回文数，个位与万位相同，十位与千位相同。
func main() {
	var result bool = true
	var x = 12343212
	var max = intLen(x)
	for i := 0; i < max/2; i++ {
		h := x % int(math.Pow10(max-i)) / int(math.Pow10(max-i-1))
		l := x % int(math.Pow10(i+1)) / int(math.Pow10(i))
		if h != l {
			result = false
			break
		}
	}

	if result {
		fmt.Printf("%d 是一个回文数.\n", x)
	} else {
		fmt.Printf("%d 不是一个回文数.\n", x)
	}
}

func intLen(n int) int {
	return len(strconv.Itoa(n))
}
