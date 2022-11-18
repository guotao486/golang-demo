/*
 * @Author: GG
 * @Date: 2022-11-18 17:03:49
 * @LastEditTime: 2022-11-18 17:19:26
 * @LastEditors: GG
 * @Description: 分解质因数
 * @FilePath: \练手算法\pro12\main.go
 *
 */
package main

import "fmt"

// 将一个正整数分解质因数。例如：输入 90，打印出 90 = 2 * 3 * 3 * 5。

// 对 n 进行分解质因数，应先找到一个最小的质数 k，然后按下述步骤完成：
// 如果这个质数恰等于 n，则说明分解质因数的过程已经结束，打印出即可。
// 如果 n<> k，但 n 能被 k 整除，则应打印出 k 的值，并用 n 除以 k 的商,作为新的正整数 n，重复执行第一步。
// 如果 n 不能被 k 整除，则用 k+1 作为 k 的值，重复执行第一步。
func main() {
	var num int = 90
	fmt.Print("请输入数字：")
	fmt.Scanf("%d\n", &num)
	fmt.Printf("%d = ", num)
	for i := 2; i < num; i++ {

		// for 循环 能整除就一直除下去
		for num != i {
			if num%i == 0 {
				fmt.Printf("%d * ", i)
				num = num / i
			} else {
				break
			}
		}
	}
	fmt.Printf("%d\n", num)
}
