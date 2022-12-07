/*
 * @Author: GG
 * @Date: 2022-12-07 14:33:18
 * @LastEditTime: 2022-12-07 15:16:39
 * @LastEditors: GG
 * @Description: 抽象工厂模式
 * @FilePath: \设计模式\抽象工厂模式\main.go
 *
 */
package main

import (
	"fmt"
)

// ----------------------形状接口--------------------------
// Shape 形状接口
type Shape interface {
	Draw()
}

// 实现形状接口的 圆形类
type Circle struct{}

// 实现形状接口的 正方形类
type Square struct{}

// 圆形类 实现 draw方法
func (c Circle) Draw() {
	fmt.Println("画圆")
}

// 正方形类 实现 draw方法
func (s Square) Draw() {
	fmt.Println("画正方形")
}

// -----------------------形状接口end-----------------------------

// -----------------------色彩接口-----------------------
// Color 色彩接口
type Color interface {
	Fill()
}

// 实现色彩接口的 红色类
type Red struct{}

// 实现色彩接口的 绿色类
type Green struct{}

// 红色类 实现 fill方法
func (r Red) Fill() {
	fmt.Println("填充红色")
}

//绿色类 实现 fill方法
func (g Green) Fill() {
	fmt.Println("填充绿色")
}

// -------------------------色彩接口end-------------------------------

// 抽象工厂接口
type AbstractFactory interface {
	GetShape(shapeName string) Shape
	GetColor(colorName string) Color
}

// -------------------------模型工厂接口--------------------------
// 模型工厂类
type ShapeFactory struct{}

// 模型工厂实例获取模型子类的方法
func (f ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	}
	return nil
}

// 模型工厂实例不需要颜色方法
func (f ShapeFactory) GetColor(colorName string) Color {
	return nil
}

// -------------------------模型工厂接口end--------------------------

// -------------------------色彩工厂接口-----------------------------
// 色彩工厂类
type ColorFactory struct{}

// 色彩工厂实例不需要模型方法
func (f ColorFactory) GetShape(shapeName string) Shape {
	return nil
}

// 色彩工厂实例获取色彩方法
func (f ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

// -------------------------色彩工厂接口end---------------------------

// -----------------------超级工厂 -------------------------
// 超级工厂，用来获取工厂实例
type FactoryProducer struct{}

// 获取工厂方法
func (f FactoryProducer) GetFactory(factoryName string) AbstractFactory {
	switch factoryName {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

// 实例化超级工厂
func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}

// 在工厂模式中，一个具体的工厂对应一种具体的产品。但是，有时候我们需要工厂可以提供多个产品对象，而不是单一产品对象
func main() {

	superFactory := NewFactoryProducer()

	color := superFactory.GetFactory("color")
	shape := superFactory.GetFactory("shape")

	red := color.GetColor("red")
	green := color.GetColor("green")

	circle := shape.GetShape("circle")
	square := shape.GetShape("square")

	//
	circle.Draw()
	red.Fill()

	square.Draw()
	green.Fill()

	circle.Draw()
	green.Fill()

}
