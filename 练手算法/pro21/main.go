/*
 * @Author: GG
 * @Date: 2022-12-02 16:17:53
 * @LastEditTime: 2022-12-02 16:31:47
 * @LastEditors: GG
 * @Description: 菱形图案
 * @FilePath: \练手算法\pro21\main.go
 *
 */
package main

import "fmt"

//打印出菱形图案
func main() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 2-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= 2*i; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := 0; i <= 2; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= 4-2*i; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
