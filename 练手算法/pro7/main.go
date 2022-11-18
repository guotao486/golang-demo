/*
 * @Author: GG
 * @Date: 2022-11-18 10:53:36
 * @LastEditTime: 2022-11-18 10:58:37
 * @LastEditors: GG
 * @Description: 乘法口诀
 * @FilePath: \练手算法\pro7\main.go
 *
 */
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
