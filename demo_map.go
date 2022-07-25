/*
 * @Author: GG
 * @Date: 2022-07-25 10:29:51
 * @LastEditTime: 2022-07-25 10:42:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_map.go
 *
 */
package main

import "fmt"

func main1() {
	var a map[string]string //定义
	fmt.Printf("a: %T\n", a)
	fmt.Printf("a: %v\n", a)

	// 不用make会报错
	a = make(map[string]string)
	a["name"] = "tom"
	fmt.Printf("a: %v\n", a)

	// 直接make
	b := make(map[string]string)

	b["name"] = "alex"
	fmt.Printf("b: %v\n", b)

	// 定义初始值
	c := map[string]string{
		"name": "clown",
	}
	fmt.Printf("c: %v\n", c)

	// 取值
	fmt.Printf("a[\"name\"]: %v\n", a["name"])

	//判断值是否存在
	v, ok := b["age"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)

	for k, v := range c {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

	for k := range c {
		fmt.Printf("k: %v\n", k)
	}
}
