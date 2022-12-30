/*
 * @Author: GG
 * @Date: 2022-12-30 09:42:47
 * @LastEditTime: 2022-12-30 09:57:38
 * @LastEditors: GG
 * @Description: 快速排序
 * @FilePath: \数据结构和算法\快速排序\main.go
 *
 */

/* 随机选择数组中的一个数 A，以这个数为基准（一般是第一个）
其他数字跟这个数进行比较，比这个数小的放在其左边，大的放到其右边
经过一次循环之后，A 左边为小于 A 的，右边为大于 A 的
这时候将左边和右边的数再递归上面的过程 */
package main

import "fmt"

func sort(arr []int) []int {
	// 总长度
	length := len(arr)
	// 判断length是否为1，一个元素就不用排了
	if length <= 1 {
		return arr
	}

	// 假设第一个元素是中间值
	middle := arr[0]
	// 左边元素
	var left []int
	// 右边元素
	var right []int

	for i := 1; i < length; i++ {
		if middle < arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	// middle_s := []int{middle}

	left = sort(left)

	right = sort(right)

	result := append(append(left, middle), right...)
	return result
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Printf("values: %v\n", values)
	arr := sort(values)
	fmt.Printf("arr: %v\n", arr)
}
