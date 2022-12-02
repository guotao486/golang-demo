/*
 * @Author: GG
 * @Date: 2022-12-02 10:32:33
 * @LastEditTime: 2022-12-02 10:37:10
 * @LastEditors: GG
 * @Description: 最大公约数 最小公倍数
 * @FilePath: \练手算法\pro14\main.go
 *
 */
package main

import "fmt"

// 求2个数的最大公约数和最小公倍数
func main() {
	var m, n, x, r int
	fmt.Printf("请输入2个数字：")
	fmt.Scanf("%d%d", &m, &n)
	x = m * n
	for n != 0 {
		r = m % n
		m = n
		n = r
	}
	fmt.Printf("最大公约数：%d,最小公倍数：%d\n", m, x/m)
}
