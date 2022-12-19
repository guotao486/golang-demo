/*
 * @Author: GG
 * @Date: 2022-12-19 09:58:20
 * @LastEditTime: 2022-12-19 10:21:16
 * @LastEditors: GG
 * @Description: 备忘录模式
 * @FilePath: \设计模式\备忘录模式\main.go
 *
 */
package main

import "fmt"

// 备忘录类
type Memento struct {
	state string
}

// 实例化备忘录类
func NewMemento(st string) *Memento {
	return &Memento{
		state: st,
	}
}

// 获取备忘录的状态
func (m *Memento) GetState() string {
	return m.state
}

// 初始类
type Originator struct {
	state string
}

// 实例化初始类
func NewOriginator(st string) *Originator {
	return &Originator{
		state: st,
	}
}

// 初始类 设置状态
func (o *Originator) SetState(st string) {
	o.state = st
}

// 初始类 获取状态
func (o *Originator) GetState() string {
	return o.state
}

// 将初始类状态保存到备忘录
func (o *Originator) SaveStateToMemento() *Memento {
	return NewMemento(o.state)
}

// 将备忘录类的状态读取到初始类
func (o *Originator) GetStateFromMemento(m *Memento) {
	o.state = m.state
}

// 保存类，用来保存备忘录实例
type CareTaker struct {
	MementoList map[int]*Memento
}

// 实例化保存类
func NewCareTaker() *CareTaker {
	return &CareTaker{
		MementoList: make(map[int]*Memento),
	}
}

// 保存类 添加备忘录实例
func (ct *CareTaker) Add(index int, m *Memento) {
	ct.MementoList[index] = m
}

// 保存类 读取备忘录实例
func (ct *CareTaker) Get(index int) *Memento {
	return ct.MementoList[index]
}

func main() {

	careTaker := NewCareTaker()
	originator := NewOriginator("状态：#1")
	originator.SetState("状态： #2")
	careTaker.Add(1, originator.SaveStateToMemento())
	originator.SetState("状态： #3")
	careTaker.Add(2, originator.SaveStateToMemento())
	originator.SetState("状态： #4")
	careTaker.Add(3, originator.SaveStateToMemento())

	fmt.Println("当前状态：", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("第一次保存的状态：", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(2))
	fmt.Println("第二次保存的状态：", originator.GetState())
}
