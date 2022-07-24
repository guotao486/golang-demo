/*
 * @Author: GG
 * @Date: 2022-07-24 17:46:36
 * @LastEditTime: 2022-07-24 17:51:31
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_switch.go
 *
 */
package main

import "fmt"

func main1() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	}

	score := 90
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score < 90 && score > 70:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

	//fallthrough  可以同时执行下一条case
	a := 100
	switch a {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	default:
		fmt.Println("end")
	}
}
