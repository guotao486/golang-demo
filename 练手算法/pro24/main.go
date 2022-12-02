/*
 * @Author: GG
 * @Date: 2022-12-02 17:10:14
 * @LastEditTime: 2022-12-02 17:13:02
 * @LastEditors: GG
 * @Description: 递归求阶乘
 * @FilePath: \练手算法\pro24\main.go
 *
 */
package main

import "fmt"

// 利用递归方法求 5!。
// fn = fn(1)*4!。
func main() {
	fmt.Println("5! 的阶乘：", face(5))
}
func face(n int) (sum int) {
	if n == 0 {
		return 1
	}
	sum = n * face(n-1)
	return
}
