/*
 * @Author: GG
 * @Date: 2022-08-12 14:14:04
 * @LastEditTime: 2022-08-12 14:33:47
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\builtin\demo_builtin.go
 *
 */
package main

import "fmt"

// append 给切片添加元素
func test1() {
	s1 := []int{1, 2, 3}
	i := append(s1, 4)
	fmt.Printf("i: %v\n", i)

	s2 := []int{7, 8, 9}
	i2 := append(s1, s2...)
	fmt.Printf("i2: %v\n", i2)
}

// len 返回数组切片字符串通道的长度
func test2() {
	s1 := "hello world"
	i := len(s1)
	fmt.Printf("i: %v\n", i)

	s2 := []int{1, 2, 3}
	fmt.Printf("len(s2): %v\n", len(s2))
}

func test3() {
	b := new(bool)
	fmt.Printf("b: %v\n", *b)
}
func main() {
	// test1()
	test2()

	print("hello world")
}
