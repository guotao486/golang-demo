/*
 * @Author: GG
 * @Date: 2022-12-29 16:35:59
 * @LastEditTime: 2022-12-29 16:49:35
 * @LastEditors: GG
 * @Description: 选择排序
 * @FilePath: \数据结构和算法\选择排序\main.go
 *
 */
package main

import "fmt"

/* 将第一个值看成最小值

然后和后续的比较找出最小值和下标

交换本次遍历的起始值和最小值

说明：每次遍历的时候，将前面找出的最小值，看成一个有序的列表，后面的看成无序的列表，然后每次遍历无序列表找出最小值。 */

// 升序
func Asort(values []int) {
	// 总循环次数
	for i := 0; i < len(values)-1; i++ {
		// 假设最小索引第一个
		min_index := i
		// 从第一个开始比较
		for j := i + 1; j < len(values); j++ {
			// 如果后面的元素比当前值小
			if values[min_index] > values[j] {
				// 修改当前最小索引
				min_index = j
			}

			//交换
			temp := values[i]
			values[i] = values[min_index]
			values[min_index] = temp
		}
	}
	fmt.Printf("values: %v\n", values)
}

// 降序
func Zsort(values []int) {
	for i := 0; i < len(values); i++ {
		max_index := i
		for j := i + 1; j < len(values); j++ {
			if values[max_index] < values[j] {
				max_index = j
			}

			temp := values[i]
			values[i] = values[max_index]
			values[max_index] = temp
		}
	}
	fmt.Printf("values: %v\n", values)
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	Asort(values)
	Zsort(values)
}
