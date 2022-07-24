/*
 * @Author: GG
 * @Date: 2022-07-24 16:02:37
 * @LastEditTime: 2022-07-24 16:08:58
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_bool.go
 *
 */
package main

import "fmt"

func main1() {
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false

	b5 := true
	b6 := false

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)

	//无法使用0 or !0表示真假
	//条件判断
	age := 17
	ok := age >= 18
	if ok {
		fmt.Println("已经成年")
	} else {
		fmt.Println("未成年")
	}

	//循环
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}

	//逻辑表达式
	age = 18
	gender := "男"

	if age >= 18 && gender == "男" {
		fmt.Println("成年男性")
	}

}
