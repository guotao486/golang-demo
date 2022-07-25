/*
 * @Author: GG
 * @Date: 2022-07-25 17:53:09
 * @LastEditTime: 2022-07-25 17:53:28
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_func_type.go
 *
 */
package main

import (
	"fmt"
)

func main1() {
	NewClient().Say()
}

//定义一个结构体
type Client struct {
	cmdable
}

//定义一个函数类型
type cmdable func(str string) error

//初始化操作
func NewClient() *Client {
	c := Client{}
	c.cmdable = SayChinese
	return &c
}

//具体逻辑的函数
func SayChinese(str string) error {
	/*
		do something
	*/
	fmt.Println("输出内容 " + str)
	return nil
}

//接收者是一个函数类型，实际的执行逻辑是
func (c cmdable) Say() {
	str := "测试"
	//调用c()函数本体
	c(str)
}
