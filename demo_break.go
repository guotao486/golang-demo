/*
 * @Author: GG
 * @Date: 2022-07-24 22:10:08
 * @LastEditTime: 2022-07-24 22:15:53
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_break.go
 *
 */
package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("i: %v\n", i)
	}
}
func test2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
		fallthrough
	case 3:
		fmt.Println("3")
		break
	default:
		fmt.Println("default")
		break

	}
}
func test3() {
LABEL:
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			if i == 5 {
				break LABEL
			}
			fmt.Printf("i: %v\n", i)
		}
		fmt.Println("end2 ...")
	}
	fmt.Println("end1 ...")
}
func main1() {
	test3()
}
