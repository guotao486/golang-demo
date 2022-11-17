/*
 * @Author: GG
 * @Date: 2022-11-17 16:30:12
 * @LastEditTime: 2022-11-17 17:10:28
 * @LastEditors: GG
 * @Description: 数字排序
 * @FilePath: \练手算法\pro5\main.go
 *
 */
package main

import "fmt"

// 输入三个 整数 x，y，z，请把这三个数由小到大输出。
func main() {
	test1()
	test2()
}

func test1() {
	var x, y, z int = 0, 0, 0
	fmt.Print("请输入：")
	fmt.Scanf("%d%d%d", &x, &y, &z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}

	if y > z {
		y, z = z, y
	}

	fmt.Printf("%d < %d < %d", x, y, z)
}

func test2() {
	var arr []int
	arr = []int{3, 23, 1, 5, 6, 17, 8, 28}
	fmt.Printf("arr: %v\n", arr)
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr); i++ {
			if (i + 1) == len(arr) {
				break
			}
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	fmt.Printf("arr: %v\n", arr)
}
