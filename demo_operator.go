/*
 * @Author: GG
 * @Date: 2022-07-24 16:47:10
 * @LastEditTime: 2022-07-24 17:15:37
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_operator.go
 *
 */
package main

import "fmt"

func main1() {
	// 算数运算符
	a := 100
	b := 10

	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a %% b): %v\n", (a % b))

	//自增
	a++
	fmt.Printf("a: %v\n", a)
	//自减
	b--
	fmt.Printf("b: %v\n", b)

	//关系运算符
	a = 1
	b = 2

	fmt.Printf("(a > b): %v\n", (a > b))
	fmt.Printf("(a < b): %v\n", (a < b))
	fmt.Printf("(a >= b): %v\n", (a >= b))
	fmt.Printf("(a <= b): %v\n", (a <= b))
	fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a != b): %v\n", (a != b))

	//逻辑运算符
	c := true
	d := false

	fmt.Printf("(c && d): %v\n", (c && d))
	fmt.Printf("(c || d): %v\n", (c || d))
	fmt.Printf("(!c): %v\n", (!c))
	fmt.Printf("(!d): %v\n", (!d))

	//位运算符
	a = 4 // 二进制 100
	fmt.Printf("a: %b\n", a)
	b = 8 // 二进制 1000
	fmt.Printf("b: %b\n", b)

	fmt.Printf("(a & b): %v, %b \n", (a & b), (a & b))
	fmt.Printf("(a | b): %v, %b\n", (a | b), (a | b))
	fmt.Printf("(a ^ b): %v, %b\n", (a ^ b), (a ^ b))
	fmt.Printf("(a << 2): %v, %b\n", (a << 2), (a << 2))
	fmt.Printf("(b >> 2): %v, %b\n", (b >> 2), (b >> 2))

	//赋值运算符
	a = 100
	fmt.Printf("a: %v\n", a)
	a += 1 // a = a + 1
	fmt.Printf("a: %v\n", a)
	a -= 1 // a = a -1
	fmt.Printf("a: %v\n", a)
	a *= 2 // a = a * 2
	fmt.Printf("a: %v\n", a)
	a /= 2 // a = a / 2
	fmt.Printf("a: %v\n", a)
}
