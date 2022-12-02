/*
 * @Author: GG
 * @Date: 2022-12-02 14:36:02
 * @LastEditTime: 2022-12-02 14:45:38
 * @LastEditors: GG
 * @Description: 计算自由落体总共的高度
 * @FilePath: \练手算法\pro18\main.go
 *
 */
package main

import "fmt"

// 一球从100 米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第 10 次落地时，共经过多少米？第 10 次反弹多高？
func main() {
	s := 100.0
	h := s / 2
	fmt.Printf("s: %v\n", s)
	fmt.Printf("h: %v\n", h)
	for i := 2; i <= 10; i++ {
		s += 2 * h // 每次回弹距离，上次的一半 * 2
		fmt.Printf("s: %v\n", s)
		h /= 2 // 反弹距离
		fmt.Printf("h: %v\n", h)
	}

}
