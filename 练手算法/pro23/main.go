/*
 * @Author: GG
 * @Date: 2022-12-02 16:53:07
 * @LastEditTime: 2022-12-02 17:01:11
 * @LastEditors: GG
 * @Description: 求阶乘和
 * @FilePath: \练手算法\pro23\main.go
 *
 */
package main

import "fmt"

//求 1+2!+3!+…+20! 的和。 !代表阶乘
func main() {
	s, t := 0, 1
	for n := 1; n <= 20; n++ {
		t *= n
		s += t
	}
	fmt.Printf("t: %v\n", s)
}
