/*
 * @Author: GG
 * @Date: 2022-07-24 17:52:24
 * @LastEditTime: 2022-07-24 18:09:28
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_for.go
 *
 */
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	a := 5
	for ; a < 10; a++ {
		fmt.Printf("a: %v\n", a)
	}
	b := 0
	for b <= 10 {
		fmt.Printf("b: %v\n", b)
		b++
	}

	var c = [5]int{1, 2, 3, 4, 5}
	for i, v := range c {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}

	var d = []int{1, 2, 3, 4, 5}
	for i, v := range d {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}

	var e = make(map[string]string)
	e["name"] = "tom"
	fmt.Printf("e: %v\n", e)
	for i, v := range e {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}

}
