/*
 * @Author: GG
 * @Date: 2022-07-25 20:57:19
 * @LastEditTime: 2022-07-25 21:01:05
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_type.go
 *
 */
package main

import "fmt"

func main1() {
	// 自定义类型定义
	type MyInt int
	// i 为MyInt类型
	var i MyInt
	i = 100
	fmt.Printf("i: %v i: %T\n", i, i)

	// 类型别名定义
	type MyInt2 = int
	// j 其实还是int类型
	var j MyInt2
	j = 100
	fmt.Printf("j: %v j:%T\n", j, j)
}
