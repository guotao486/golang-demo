/*
 * @Author: GG
 * @Date: 2022-07-25 17:43:05
 * @LastEditTime: 2022-07-25 20:56:46
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_needle.go
 *
 */
package main

import "fmt"

func main1() {
	var a = 10
	ip := &a
	fmt.Printf("a: %p\n", &a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip)

	var ip2 *int
	ip2 = &a
	fmt.Printf("ip2: %v\n", ip2)
	fmt.Printf("ip2: %v\n", *ip2)

	// 数组指针
	var i int
	var ptr [3]*int
	fmt.Printf("ptr: %v\n", ptr)

	b := []int{1, 3, 5}
	fmt.Printf("b: %p\n", b)
	fmt.Printf("b: %v\n", b)
	for i = 0; i < 3; i++ {
		ptr[i] = &b[i]
	}
	fmt.Printf("ptr: %v\n", ptr)
	for i := 0; i < 3; i++ {
		fmt.Printf("ptr: %v\n", *ptr[i])
	}

}
