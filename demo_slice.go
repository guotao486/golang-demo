/*
 * @Author: GG
 * @Date: 2022-07-25 09:52:36
 * @LastEditTime: 2022-07-25 22:04:15
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_slice.go
 *
 */
package main

import "fmt"

func main1() {
	var a []int
	fmt.Printf("a: %v\n", a)

	var b = []int{1, 2, 3, 4, 5}
	fmt.Printf("b: %v\n", b)

	var c = make([]int, 5)
	c[0] = 1
	c[1] = 2
	c[2] = 3
	c[3] = 4
	c[4] = 5

	fmt.Printf("c: %v\n", c)
	fmt.Printf("c[:]: %v\n", c[:])
	fmt.Printf("c[1]: %v\n", c[1])
	fmt.Printf("c[0:]: %v\n", c[0:])
	fmt.Printf("c[0:2]: %v\n", c[0:2])
	fmt.Printf("c[2:4]: %v\n", c[2:4])
	fmt.Printf("c[:3]: %v\n", c[:3])

	var d []int
	d = append(d, 1)
	d = append(d, 2)
	fmt.Printf("d: %v\n", d)

	e := []int{}
	e = append(e, 1)
	e = append(e, 2)
	fmt.Printf("e: %v\n", e)

	// 合并切片
	f := append(d, e...)
	fmt.Printf("f: %v\n", f)

	//删除切片
	g := []int{1, 2, 3, 4, 5}
	g = append(g[:3], g[4:]...)
	fmt.Printf("g: %v\n", g)

	// 浅拷贝，指向同一个内存地址，一个值改变另一个也会改变
	h := []int{5, 4, 3, 2, 1}
	fmt.Printf("h: %p\n", h)
	fmt.Printf("h: %v\n", h)

	i := h
	fmt.Printf("i: %p\n", i)
	i[0] = 100
	fmt.Printf("i: %v\n", i)
	fmt.Printf("h: %v\n", h)

	// 深拷贝，拷贝出一个独立的内存地址和值，互不影响
	j := make([]int, len(i))
	copy(j, h)
	fmt.Printf("j: %p\n", j)
	fmt.Printf("j: %v\n", j)

	j[0] = 500
	fmt.Printf("j: %v\n", j)
	fmt.Printf("h: %v\n", h)
	fmt.Printf("i: %v\n", i)
}
