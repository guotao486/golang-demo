/*
 * @Author: GG
 * @Date: 2022-11-18 16:02:22
 * @LastEditTime: 2022-11-18 16:18:52
 * @LastEditors: GG
 * @Description: 找出素数
 * @FilePath: \练手算法\pro10\main.go
 *
 */
package main

import "fmt"

// 素数，除了1和该数自身外，无法被其他自然数整除的数，也就是从2开始到（自身-1）能被整除 就不是素数
func main() {
	count := 0
	isPrime := func(n int) bool {
		for j := 2; j < n; j++ {
			if n%j == 0 {
				return false
			}
		}
		return true
	}

	for i := 1; i < 200; i++ {
		b := isPrime(i)
		if b {
			fmt.Printf("i:%d\n", i)
			count++
		}
	}

	fmt.Printf("count: %v\n", count)
}
