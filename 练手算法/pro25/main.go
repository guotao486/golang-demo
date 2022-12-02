/*
 * @Author: GG
 * @Date: 2022-12-02 17:17:23
 * @LastEditTime: 2022-12-02 17:23:47
 * @LastEditors: GG
 * @Description: 倒叙输出 字符
 * @FilePath: \练手算法\pro25\main.go
 *
 */
package main

import "fmt"

// 利用递归 函数 调用方式，将所输入的 `5`个字符，以相反顺序打印出来。
func main() {
	fmt.Printf("请输入%d个字符", 5)

	putchar(5)
}

func putchar(n int) {
	var c byte
	if n >= 1 {
		fmt.Scanf("%c", &c) // 最外层的最先输入
		putchar(n - 1)      // 先执行完内部的putchar
		fmt.Printf("%c", c) // 最外层的最后输入达到倒序输出
		fmt.Println("")
	}
}
