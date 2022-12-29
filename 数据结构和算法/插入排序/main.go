/*
 * @Author: GG
 * @Date: 2022-12-29 17:13:35
 * @LastEditTime: 2022-12-29 17:17:23
 * @LastEditors: GG
 * @Description: 插入排序
 * @FilePath: \数据结构和算法\插入排序\main.go
 *
 */

/* 默认从第二个数据开始比较。

如果第二个数据比第一个小，则交换。然后在用第三个数据比较，如果比前面小，则插入。否则，退出循环

说明：默认将第一数据看成有序列表，后面无序的列表循环每一个数据，如果比前面的数据小则插入. */
package main

import "fmt"

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Printf("values: %v\n", values)
	// values: [93 4 84 85 80 37 81 93 27 12]
	// values: [93 84 4 85 80 37 81 93 27 12]
	// values: [93 84 85 4 80 37 81 93 27 12]
	// values: [93 85 84 4 80 37 81 93 27 12]
	// values: [93 85 84 80 4 37 81 93 27 12]
	// values: [93 85 84 80 37 4 81 93 27 12]
	// values: [93 85 84 80 37 81 4 93 27 12]
	// values: [93 85 84 80 81 37 4 93 27 12]
	// values: [93 85 84 81 80 37 4 93 27 12]
	// values: [93 85 84 81 80 37 93 4 27 12]
	// values: [93 85 84 81 80 93 37 4 27 12]
	// values: [93 85 84 81 93 80 37 4 27 12]
	// values: [93 85 84 93 81 80 37 4 27 12]
	// values: [93 85 93 84 81 80 37 4 27 12]
	// values: [93 93 85 84 81 80 37 4 27 12]
	// values: [93 93 85 84 81 80 37 27 4 12]
	// values: [93 93 85 84 81 80 37 27 12 4]

	// 降序
	for i := 0; i < len(values); i++ {
		for j := i; j > 0; j-- {
			if values[j] > values[j-1] {
				temp := values[j]
				values[j] = values[j-1]
				values[j-1] = temp
				fmt.Printf("values: %v\n", values)
			}
		}
	}

	fmt.Printf("values: %v\n", values)
}
