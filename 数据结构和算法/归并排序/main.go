/*
 * @Author: GG
 * @Date: 2022-12-30 10:01:35
 * @LastEditTime: 2022-12-30 11:11:43
 * @LastEditors: GG
 * @Description: 归并排序
 * @FilePath: \数据结构和算法\归并排序\main.go
 *
 */

/* 首先将数组一份为二，分别为左数组和右数组
若左数组的长度大于1，那么对左数组实施归并排序
若右数组的长度大于1， 那么对右数组实施归并排序
将左右数组进行合并 */

package main

import "fmt"

/**
 * @description:
 * @param {[]int} arr
 * @param {*} a 起始
 * @param {int} b 结束
 * @return {*}
 */
func sort(arr []int, a, b int) {
	// 长度小于等于1，不用排序
	if b-a <= 1 {
		return
	}

	// 取中间值
	c := (a + b) / 2
	//递归调用
	// 左边
	sort(arr, a, c)
	arrLeft := make([]int, c-a)
	copy(arrLeft, arr[a:c])
	// 右边
	sort(arr, c, b)
	arrRight := make([]int, b-c)
	copy(arrRight, arr[c:b])

	i := 0
	j := 0

	for k := a; k < b; k++ {
		if i >= c-a {
			arr[k] = arrRight[j]
			j++
		} else if j >= b-c {
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}

}

func main() {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(arr)
	sort(arr, 0, len(arr))
	fmt.Println(arr)
	// [4 93 84 85 80 37 81 93 27 12]
	// [4 12 27 37 80 81 84 85 93 93]
}
