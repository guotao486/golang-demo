/*
 * @Author: GG
 * @Date: 2022-12-09 10:31:32
 * @LastEditTime: 2022-12-09 10:50:43
 * @LastEditors: GG
 * @Description: 享元模式
 * @FilePath: \设计模式\享元模式\main.go
 *
 */
package main

import "fmt"

// 模型接口
type Shape interface {
	Draw()
}

// 圆形类
type Circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

// 实例化圆形类
func NewCircle(color string) *Circle {
	return &Circle{
		Color: color,
	}
}

// 设置x
func (c Circle) SetX(x int) {
	c.X = x
}

// 设置Y
func (c Circle) SetY(y int) {
	c.Y = y
}

// 设置radius
func (c Circle) SetRadius(radius int) {
	c.Radius = radius
}

// 圆形类实现接口Draw方法
func (c Circle) Draw() {
	fmt.Printf("绘制圆 [颜色: %s, x: %d, y: %d, 半径: %d] \n",
		c.Color,
		c.X,
		c.Y,
		c.Radius)
}

// 模型工厂类，包含一个circle的map
type ShapeFactory struct {
	circleMap map[string]Shape
}

// 实例化模型工厂
func NewShapeCirlce() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]Shape),
	}
}

/**
 * @description: 获取一个圆形实例
 * @param {string} color 实例的颜色
 * @return {*}
 */
func (c ShapeFactory) GetCircle(color string) Shape {
	// 从map中找出
	circle := c.circleMap[color]
	// 没有就创建
	if circle == nil {
		circle = NewCircle(color)
		c.circleMap[color] = circle
		fmt.Println("创建圆的颜色", color)
	}
	return circle
}
func main() {
	sf := NewShapeCirlce()
	sf.GetCircle("red").Draw()
	sf.GetCircle("green").Draw()
	sf.GetCircle("red").Draw()
}
