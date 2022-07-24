/*
 * @Author: GG
 * @Date: 2022-07-24 17:15:46
 * @LastEditTime: 2022-07-24 17:46:29
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_if.go
 *
 */
package main

import "fmt"

func main1() {
	var flag = true
	if flag {
		fmt.Println("flag is true")
	}
	fmt.Println("flag end....")

	var age = 19
	if age > 20 {
		fmt.Println("age > 19")
	}
	fmt.Println("age end...")

	if age = 20; age > 20 {
		fmt.Println("age > 20")
	}
	fmt.Println("age end...")

	// if else
	a := 1
	b := 2
	if a > b {
		fmt.Printf("(a > b): %v\n", (a > b))
	} else {
		fmt.Printf("(a <= b): %v\n", (a <= b))
	}

	// if else if
	score := 81
	if score < 60 {
		fmt.Println("C")
	} else if score >= 60 && score <= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
