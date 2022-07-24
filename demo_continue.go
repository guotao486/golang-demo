/*
 * @Author: GG
 * @Date: 2022-07-24 22:23:24
 * @LastEditTime: 2022-07-24 22:27:37
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_continue.go
 *
 */
package main

import "fmt"

func main1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i: %v\n", i)
	}

	for i := 0; i < 5; i++ {
	MY_LABEL:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue MY_LABEL
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}
}
