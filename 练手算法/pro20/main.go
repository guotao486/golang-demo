/*
 * @Author: GG
 * @Date: 2022-12-02 15:32:31
 * @LastEditTime: 2022-12-02 16:07:50
 * @LastEditors: GG
 * @Description: 找出对手
 * @FilePath: \练手算法\pro20\main.go
 *
 */
package main

import "fmt"

// 两个乒乓球队进行比赛，各出三人。甲队为 a、b、c 三人，乙队为x、y、z 三人。
// 已抽签决定比赛名单。有人向队员打听比赛的名单。a 说他不和 x 比，c 说他不和 x、z 比，请编程序找出三队赛手的名单。
func main() {
	var B = []string{"x", "y", "z"}
	var i, j, k string // a b c

	for _, i = range B {
		for _, j = range B {
			if i != j {
				for _, k = range B {
					if i != k && j != k {
						// a 说他不和 x 比
						if i != "x" && k != "x" && k != "z" {
							fmt.Printf("a -- %s\nb -- %s\nc -- %s", i, j, k)

							break
						}
					}
				}
			}

		}
	}

}
