/*
 * @Author: GG
 * @Date: 2022-12-02 11:11:21
 * @LastEditTime: 2022-12-02 11:25:03
 * @LastEditors: GG
 * @Description: 统计一个 字符串 中各个字符的个数
 * @FilePath: \练手算法\pro15\main.go
 *
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

//输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数
// 你用rune 去进行比较统计
func main() {
	var i, j, k, l int = 0, 0, 0, 0
	fmt.Printf("请输入一串字符：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, ch := range input {
		switch {
		case ch >= 'A' && ch <= 'Z':
			i++
		case ch >= 'a' && ch <= 'z':
			i++
		case ch == ' ' || ch == '\t':
			j++
		case ch >= '0' && ch <= '9':
			k++
		default:
			l++
		}
	}
	fmt.Printf("char count = %d, space count = %d, digit count = %d, others count = %d", i, j, k, l)
}
