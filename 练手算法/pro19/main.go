/*
 * @Author: GG
 * @Date: 2022-12-02 15:17:03
 * @LastEditTime: 2022-12-02 15:26:20
 * @LastEditors: GG
 * @Description: 猴子吃桃
 * @FilePath: \练手算法\pro19\main.go
 *
 */
package main

import "fmt"

// 猴子吃桃问题：猴子第一天摘下若干个桃子，当即吃了一半，还不瘾，又多吃了一个。
// 第二天早上又将剩下的桃子吃掉一半，又多吃了一个。
// 以后每天早上都吃了前一天剩下的一半零一个。
// 到第 10 天早上想再吃时，见只剩下一个桃子了。求第一天共摘了多少。
func main() {
	var x1, x2, day int = 0, 1, 9
	for day > 0 {
		x1 = (x2 + 1) * 2
		x2 = x1
		day--
	}
	fmt.Printf("x1: %v\n", x1)
}
