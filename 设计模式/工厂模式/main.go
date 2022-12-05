/*
 * @Author: GG
 * @Date: 2022-12-05 16:00:56
 * @LastEditTime: 2022-12-05 16:19:48
 * @LastEditors: GG
 * @Description:
 * @FilePath: \设计模式\工厂模式\main.go
 *
 */
package main

import "fmt"

// 水果接口
type Fruit interface {
	// 种植水果
	grant()

	// 采摘水果
	pick()
}

type Apple struct {
}

func (a *Apple) grant() {
	fmt.Println("种植 apple")
}

func (a *Apple) pick() {
	fmt.Println("采摘 apple")
}

type Orange struct {
}

func (o *Orange) grant() {
	fmt.Println("种植 orange")
}

func (o *Orange) pick() {
	fmt.Println("采摘 orange")
}

// 工厂初始化入口
func NewFruit(class string) Fruit {
	switch class {
	case "apple":
		return &Apple{}
	case "orange":
		return &Orange{}
	}
	return nil
}
func main() {
	apple := NewFruit("apple")
	apple.grant()
	apple.pick()
	fmt.Println("-------------")

	orange := NewFruit("orange")
	orange.grant()
	orange.pick()

	fmt.Println("-------------")
	// banner := NewFruit("banner")
	// banner.grant()

	// 类型断言
	v, err := apple.(Fruit)
	fmt.Printf("v: %v\n", v)     //&{}
	fmt.Printf("err: %v\n", err) //rue
	v.pick()                     // 采摘 apple
}
