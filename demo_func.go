/*
 * @Author: GG
 * @Date: 2022-07-25 11:12:31
 * @LastEditTime: 2022-07-25 16:41:58
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_func.go
 *
 */
package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// 指定返回参数时，return 可为空
func compare(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func ff1() {
	fmt.Println("func f1...")
}

func ff2() (name string, age int) {
	name = "tom"
	age = 18
	return name, age
}
func ff3() (name string, age int) {
	name = "tom"
	age = 18
	return
}
func ff4() (name string, age int) {
	n := "tom"
	a := 18
	return n, a
}

func ff5(a ...int) {
	fmt.Printf("a: %v\n", a)
}

// 引用类型的值修改会影响到外部 slice map interface channel
func ff6(a []int) {
	a[0] = 100
}
func main1() {
	a := 1
	b := 2
	sum := sum(a, b)
	max := compare(a, b)
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("max: %v\n", max)

	ff5(1, 2, 3, 4)

	c := []int{1, 2, 3, 4, 5}
	fmt.Printf("c: %v\n", c)
	ff6(c)
	fmt.Printf("c: %v\n", c)
}
