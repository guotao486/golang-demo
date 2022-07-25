/*
 * @Author: GG
 * @Date: 2022-07-25 09:35:51
 * @LastEditTime: 2022-07-25 09:52:27
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_array.go
 *
 */
package main

import "fmt"

func main1() {
	// 定义数组,默认值 0 “” false
	var a [3]int
	var b [3]string
	var c [3]bool

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	//初始化
	var d = [3]int{1, 2, 3}
	var e = [3]string{"a", "b", "c"}
	var f = [3]bool{true, false, true}
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)

	//省略长度
	var g = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("g: %v\n", g)

	fmt.Printf("len(g): %v\n", len(g))
	fmt.Printf("cap(g): %v\n", cap(g))
	fmt.Printf("g[0]: %v\n", g[0])
	for i := 0; i < len(g); i++ {
		fmt.Printf("i: %v\n", i)
	}

	for i, v := range g {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}
