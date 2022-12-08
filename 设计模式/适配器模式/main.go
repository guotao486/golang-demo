/*
 * @Author: GG
 * @Date: 2022-12-08 14:29:24
 * @LastEditTime: 2022-12-08 14:34:07
 * @LastEditors: GG
 * @Description: 适配器模式
 * @FilePath: \设计模式\适配器模式\main.go
 *
 */
package main

import "fmt"

type OldInterface interface {
	OldMethod()
}

type OldImpl struct{}

func (o OldImpl) OldMethod() {
	fmt.Println("旧方法实现")
}

type NewInterface interface {
	NewMethod()
}

// 适配器结构体
type Adapter struct {
	OldInterface // 旧接口
}

// 适配器方法
func (a Adapter) NewMethod() {
	fmt.Println("新方法实现")
	a.OldMethod()
}

func main() {
	oldInterface := OldImpl{}
	a := Adapter{OldInterface: oldInterface}
	a.NewMethod()
}
