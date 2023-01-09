/*
 * @Author: GG
 * @Date: 2023-01-09 11:34:30
 * @LastEditTime: 2023-01-09 11:49:39
 * @LastEditors: GG
 * @Description: stack 栈
 * @FilePath: \数据结构和算法\栈stack\main.go
 *
 */

/*  stack的特点
是一种后进先出的数据结构，Last In First Out(LIFO)。
函数调用使用的就是一种栈结构，调用时入栈，返回时出栈。 */
package main

import "fmt"

// 利用字符串切片
type Stack []string

// 判断栈
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 入栈，后进先出，所以从尾部添加
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// 出栈,后进先出
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index] // 获得顶部元素
		*s = (*s)[:index]      // 删除顶部元素
		return element, true   // 返回
	}
}

func main() {
	var stack Stack // 创建一个栈

	stack.Push("PHP")
	stack.Push("golang")
	stack.Push("java")

	for len(stack) > 0 {
		s, _ := stack.Pop()
		fmt.Printf("s: %v\n", s)
	}
}
