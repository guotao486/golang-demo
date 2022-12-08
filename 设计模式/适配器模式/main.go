/*
 * @Author: GG
 * @Date: 2022-12-08 14:29:24
 * @LastEditTime: 2022-12-08 15:04:42
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

// 适配器模式(Adapter Pattern)是作为两个不兼容接口之间的桥梁。适配器模式将一个类的接口转换为另一个类的接口，使得原本由于接口不兼容而不能一起工作的类可以一起工作。
func main() {
	oldInterface := OldImpl{}
	a := Adapter{OldInterface: oldInterface}
	a.NewMethod()
}
