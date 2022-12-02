/*
 * @Author: GG
 * @Date: 2022-12-02 17:26:49
 * @LastEditTime: 2022-12-02 17:46:07
 * @LastEditors: GG
 * @Description: 递归推断年纪
 * @FilePath: \练手算法\pro26\main.go
 *
 */
package main

import "fmt"

// 有 5 个人坐在一起，问第五个人多少岁？他说比第 4 个人大 2 岁。
// 问第 4 个人岁数，他说比第 3 个人大 2 岁。
// 问第三个人，又说比第 2 人大两岁。
// 问第 2 个人，说比第一个人大两岁。
// 最后问第一个人，他说是 10 岁。
// 请问第五个人多大？
func main() {
	fmt.Printf("第五个人：%d岁", calcAge(5))
}

func calcAge(n int) (age int) {
	if n == 1 {
		age = 10
	} else {
		age = calcAge(n-1) + 2
	}
	return
}
