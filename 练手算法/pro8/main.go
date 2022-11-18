/*
 * @Author: GG
 * @Date: 2022-11-18 14:59:17
 * @LastEditTime: 2022-11-18 15:35:07
 * @LastEditors: GG
 * @Description: 输出国际象棋棋盘
 * @FilePath: \练手算法\pro8\main.go
 *
 */
package main

import "fmt"

// 输出国际象棋棋盘。
func main() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("□")
			} else {
				fmt.Print("■")
			}
		}
		fmt.Printf("\n")
	}
}
