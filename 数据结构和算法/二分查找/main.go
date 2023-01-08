/*
 * @Author: GG
 * @Date: 2022-12-30 11:32:22
 * @LastEditTime: 2023-01-08 11:44:12
 * @LastEditors: GG
 * @Description: 二分查找
 * @FilePath: \数据结构和算法\二分查找\main.go
 *
 */
package main

import (
	"fmt"
	"sort"
)

// 升序排序数组中搜索某个值得下标
func search(s []int, k int) int {
	//低位高位，长度
	lo, hi := 0, len(s)-1

	// 如果低位 小于等于高位一直循环
	for lo <= hi {
		// 取中间位，向右移动一位等于除以2
		m := (lo + hi) >> 1
		if s[m] < k {
			// 左边区域最后一位 小于 查询值，则左边不符合
			lo = m + 1
		} else if s[m] > k {
			// 左边区域最后一位 大于 查询值，右边不符合，查询值在左边
			hi = m - 1
		} else {
			// 不大于 不小于 代表正好是中介值
			return m
		}
	}
	return -1
}

// 二分 递归调用
func search2(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	// 退出条件，判断条件不能为等号，因为在相等时仍然要进行一次判断。
	// 退出条件为等号可能会造成明明数组中有该值，却错过判断的情况发生。
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}
	middleIndex := (leftIndex + rightIndex)

	if findVal > (*arr)[middleIndex] {
		// 查找值在右边区域
		search2(arr, middleIndex+1, rightIndex, findVal)
	} else if findVal < (*arr)[middleIndex] {
		// 查找值在左边
		search2(arr, leftIndex, middleIndex-1, findVal)
	} else {
		fmt.Println("找到了,下标：", middleIndex)
	}
}
func main() {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	findVal := 81
	sort.Ints(arr)
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("search(arr, %v): %v\n", findVal, search(arr, findVal))
	search2(&arr, 0, len(arr)-1, findVal)
}
