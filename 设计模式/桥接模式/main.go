/*
 * @Author: GG
 * @Date: 2022-12-08 15:16:41
 * @LastEditTime: 2022-12-08 15:38:12
 * @LastEditors: GG
 * @Description: 桥接模式
 * @FilePath: \设计模式\桥接模式\main.go
 *
 */
package main

import "fmt"

/*
 * 桥接是用于把抽象化与实现化解偶，使得二者可以独立变化。这种类型的设计模式属于结构型模式，它通过提供抽象化和实现化之间的桥接结构，来实现二者的解偶。
 * 这种模式涉及到一个作为桥接的接口，使得实体类的功能独立于接口实现类。这两种类型的类可被结构化改变而互不影响

 * 我们使用DrawAPI作为桥接模式的抽象接口，ShapeCirlce作为桥接模式的实体类。将抽象接口保存在实体类中，使得抽象接口实例变化时，ShapeCircle可以始终不变。
 */

//画图抽象接口，桥接模式的抽象接口
type DrawAPI interface {
	DrawCirlce(radius, x, y int)
}

// 红色圆的实体类，桥接模式接口
type RedCirlce struct{}

// 实例化红色圆
func NewRedCirlce() *RedCirlce {
	return &RedCirlce{}
}

// 红色圆实现接口方法
func (r RedCirlce) DrawCirlce(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

// -------------------------------------
// 绿色圆的实体类，桥接模式接口
type GreenCirlce struct{}

// 实例化绿色圆
func NewGreenCirlce() *GreenCirlce {
	return &GreenCirlce{}
}

// 绿色圆实现接口方法
func (g GreenCirlce) DrawCirlce(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: green, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

// ------------------------------------
// 桥接模式实体类
type ShapeCirlce struct {
	Radius  int
	X       int
	Y       int
	drawAPI DrawAPI
}

// 实例化桥接模式实体类
func NewShapeCirlce(radius, x, y int, drawAPI DrawAPI) *ShapeCirlce {
	return &ShapeCirlce{
		Radius:  radius,
		X:       x,
		Y:       y,
		drawAPI: drawAPI,
	}
}

// 实体类的draw方法，实现内部抽象类的DrawCirlce方法
func (s ShapeCirlce) Draw() {
	s.drawAPI.DrawCirlce(s.Radius, s.X, s.Y)
}

func main() {
	redCirlce := NewShapeCirlce(5, 6, 8, NewRedCirlce())
	redCirlce.Draw()

	greenCirlce := NewShapeCirlce(7, 8, 9, NewGreenCirlce())
	greenCirlce.Draw()
}
