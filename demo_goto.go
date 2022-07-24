/*
 * @Author: GG
 * @Date: 2022-07-24 22:28:44
 * @LastEditTime: 2022-07-24 22:31:23
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goto.go
 *
 */
package main

import "fmt"

func f() {
	a := 1
	if a == 1 {
		goto LABEL1
	} else {
		fmt.Println("other")
	}

LABEL1:
	fmt.Printf("next...")
}

func f2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto LABEL1
			}
			fmt.Printf("j: %v,i: %v\n", j, i)
		}
	}
LABEL1:
	fmt.Println("label1")
}
func main1() {
	f2()
}
