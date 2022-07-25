/*
 * @Author: GG
 * @Date: 2022-07-25 16:42:10
 * @LastEditTime: 2022-07-25 17:11:36
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_func2.go
 *
 */
package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hello world", name)
}

// 将函数sayHello当成参数
func ff2_1(name string, f func(string)) {
	f(name)
}

func add(x, y int) int {
	// var sum int
	// sum = x + y
	// return sum
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 返回函数
func cal(s string) func(int, int) int {
	sayHello("alx")
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}

}

// 闭包，函数内部的函数
func add2() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// base 在函数中计算，即使返回了依然计算
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 递归
func ff2_2(a int) int {
	if a < 10 {
		a++
		return ff2_2(a)
	} else {
		return a
	}
}
func main1() {
	// 将函数sayHello当成参数
	ff2_1("tom", sayHello)

	// 返回一个函数
	add := cal("+")
	r := add(1, 2)
	fmt.Printf("r: %v\n", r)
	r = add(1, 2)
	fmt.Printf("r: %v\n", r)

	fmt.Println("-----------")

	sub := cal("-")
	r = sub(100, 50)
	fmt.Printf("r: %v\n", r)

	//闭包
	a := add2()
	b := a(1)
	fmt.Printf("b: %v\n", b)
	b = a(1)
	fmt.Printf("b: %v\n", b)

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(3), f2(4)) // 12 8
	fmt.Println(f1(5), f2(6)) // 13 7

	//递归
	c := ff2_2(5)
	fmt.Printf("c: %v\n", c)
}
