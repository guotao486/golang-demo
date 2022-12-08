/*
 * @Author: GG
 * @Date: 2022-12-08 16:23:16
 * @LastEditTime: 2022-12-08 16:54:50
 * @LastEditors: GG
 * @Description: 装饰器模式
 * @FilePath: \设计模式\装饰器模式\main.go
 *
 */
package main

import "fmt"

/*
 * 装饰器模式(Decorator Pattern)允许向一个现有的对象添加新的功能，同时又不改变其结构。
 * 装饰器模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整的前提下，提供了额外的功能。
 */

// 模型接口
type Shape interface {
	Draw()
}

// 圆形类
type Circle struct{}

// 实例化圆形类
func NewCircle() *Circle {
	return &Circle{}
}

// 圆形类实现模型接口方法
func (c Circle) Draw() {
	fmt.Println("画圆方法")
}

// -------装饰器--------
// 红色装饰器
type RedShapeDecorator struct {
	DecoratorShape Shape
}

// 实例化红色装饰器
func NewRedShapeDecorator(deShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{
		DecoratorShape: deShape,
	}
}

// 装饰器方法
func (d RedShapeDecorator) SetRedBorder() {
	fmt.Println("红色边框")
}

// 实现shape接口方法
func (d RedShapeDecorator) Draw() {
	d.DecoratorShape.Draw()
	d.SetRedBorder()
}

func main() {
	c := NewCircle()
	d := NewRedShapeDecorator(c)
	d.Draw()
}
