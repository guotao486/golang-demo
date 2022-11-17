/*
 * @Author: GG
 * @Date: 2022-11-17 17:32:54
 * @LastEditTime: 2022-11-17 17:40:34
 * @LastEditors: GG
 * @Description: 打印图形
 * @FilePath: \练手算法\pro6\main.go
 *
 */
package main

import "fmt"

func main() {
	var len int = 5
	for i := 0; i < len; i++ {
		if i < 3 {
			for j := 4; j > i*2-1; j-- {
				fmt.Print("*")
			}
		} else {
			for k := 0; k < i*2-2; k++ {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}
