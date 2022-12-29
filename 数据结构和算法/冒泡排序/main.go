/*
 * @Author: GG
 * @Date: 2022-12-26 16:02:19
 * @LastEditTime: 2022-12-29 16:23:20
 * @LastEditors: GG
 * @Description: 冒泡算法
 * @FilePath: \数据结构和算法\冒泡排序\main.go
 *
 */
package main

import "fmt"

/*
冒泡排序，是通过每一次遍历获取最大/最小值

将最大值/最小值放在尾部/头部

然后除开最大值/最小值，剩下的数据在进行遍历获取最大/最小值
*/

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// 升序
func Asort(values []int) {
	// 有几个数比较几次
	for i := 0; i < len(values)-1; i++ {
		// value[i] 与 value[i+1]比较
		for j := i + 1; j < len(values); j++ {
			// 如果后面的比前面的小就交换，所以是升序
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Printf("values: %v\n", values)
}

// 降序
func Zsort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Printf("values: %v\n", values)
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Printf("values: %v\n", values)
	Asort(values)
	Zsort(values)
}
