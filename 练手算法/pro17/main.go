/*
 * @Author: GG
 * @Date: 2022-12-02 12:01:55
 * @LastEditTime: 2022-12-02 14:32:35
 * @LastEditors: GG
 * @Description: 找出所有完数
 * @FilePath: \练手算法\pro17\main.go
 *
 */
package main

import "fmt"

// 一个数如果恰好等于它的因子之和，这个数就称为 “完数”。例如 6=1＋2＋3，编程找出 1000 以内的所有完数。
func main() {
	for n := 2; n < 1000; n++ {
		m := n
		for i := 1; i < n; i++ {
			if n%i == 0 {
				m -= i
			}
		}
		if m == 0 {
			fmt.Println("num is:", n)
		}
	}
}
