/*
 * @Author: GG
 * @Date: 2022-07-24 15:01:20
 * @LastEditTime: 2022-07-24 15:09:07
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_const.go
 *
 */
package main

import "fmt"

func main1() {
	const PI float64 = 3.14
	const PI2 = 3.1415 //省略类型

	const (
		WIDTH  = 100
		HEIGHT = 200
	)

	// 多重赋值
	const i, j = 1, 2
	const a, b, c = 1, 2, "foo"

	// 省略值表示和上一行值相同
	const (
		a1 = 100
		a2
		a3
	)

	//iota 每调用一次加1 遇到const关键字重置为0
	const (
		c1 = iota
		c2 = iota
		c3 = iota
		_  //跳过一次
		c4 = iota
		c5 = 200 //跳过并插队一次
		c6 = iota
	)
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("PI2: %v\n", PI2)
	fmt.Printf("WIDTH: %v\n", WIDTH)
	fmt.Printf("HEIGHT: %v\n", HEIGHT)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("c3: %v\n", c3)
	fmt.Printf("c4: %v\n", c4)
	fmt.Printf("c5: %v\n", c5)
	fmt.Printf("c6: %v\n", c6)
}
