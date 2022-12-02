/*
 * @Author: GG
 * @Date: 2022-11-18 17:38:57
 * @LastEditTime: 2022-12-02 10:30:32
 * @LastEditors: GG
 * @Description: 三元表达式
 * @FilePath: \练手算法\pro13\main.go
 *
 */

package main

import "fmt"

type B bool

func main() {
	var score int = 0
	fmt.Printf("请输入分数：")
	fmt.Scanf("%d\n", &score)
	// 会先执行 score >= 60
	grade := B(score >= 90).C("优秀", B(score >= 60).C("及格", "不及格"))
	fmt.Printf("grade: %v\n", grade)
}

func (b B) C(t interface{}, f interface{}) interface{} {
	fmt.Printf("b: %v\n", b)
	if bool(b) == true {
		return t
	}
	return f
}
