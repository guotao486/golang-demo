/*
 * @Author: GG
 * @Date: 2022-12-02 11:47:32
 * @LastEditTime: 2022-12-02 12:00:19
 * @LastEditors: GG
 * @Description: 累计特殊数
 * @FilePath: \练手算法\pro16\main.go
 *
 */
package main

import "fmt"

//求 s = a + aa + aaa + aaaa + aa…a 的值，其中 a 是一个数字。例如 2+22+222+2222+22222(此时共有 5 个数相加)，几个数相加由键盘控制。
func main() {
	var a, n, count int = 0, 0, 0
	var sum, tn int = 0, 0
	var str string = ""
	fmt.Printf("请输入 a and n：")
	fmt.Scanf("%d%d", &a, &n)

	for count < n {
		// tn = 0 + 2  2+20
		tn = tn + a
		a = a * 10 // 20 200 2000
		sum += tn  // 2+22+222
		count++
		str += fmt.Sprintf("+%d", tn)
	}
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("str: %v\n", str)

}
