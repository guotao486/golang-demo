/*
 * @Author: GG
 * @Date: 2022-12-09 16:47:22
 * @LastEditTime: 2022-12-09 17:16:12
 * @LastEditors: GG
 * @Description: 观察者模式
 * @FilePath: \设计模式\观察者模式\main.go
 *
 */
package main

import (
	"fmt"
)

/*
 * 定义了对象之间的一对多依赖，让多个观察者对象同时监听某一个主题对象，当主题对象发生变化时，它的所有依赖者都会收到通知并更新。
 * 这种模式有时又称作发布-订阅模式、模型-视图模式，它是对象行为型模式。
 * 例如，微信公众号作者和粉丝之间的发布和订阅。
 */

// 读者（观察者）接口
type Reader interface {
	Update()
	Observer(subject *Subject)
}

type ReaderUser struct {
	Name    string
	subject *Subject
}

func NewReaderUser(name string) *ReaderUser {
	return &ReaderUser{
		Name: name,
	}
}
func (r *ReaderUser) Update() {
	fmt.Println("读者【", r.Name, "】", r.subject.GetState())
}

func (r *ReaderUser) Observer(subject *Subject) {
	r.subject = subject
	r.subject.Attach(r)
}

// 主题
type Subject struct {
	Readers []Reader // 观察者列表
	State   int
}

// 实例化
func NewSubject() *Subject {
	return &Subject{
		State:   0,
		Readers: make([]Reader, 0),
	}
}

// 添加观察则
func (s *Subject) Attach(Observer Reader) {
	s.Readers = append(s.Readers, Observer)
}

// 获得状态
func (s Subject) GetState() int {
	return s.State
}

// 修改状态
func (s *Subject) SetState(state int) {
	s.State = state
	s.NotifyAllObservers()
}

// 通知所有观察者
func (s Subject) NotifyAllObservers() {
	for _, observer := range s.Readers {
		observer.Update()
	}
}

func main() {
	subject := NewSubject()
	u1 := NewReaderUser("tom")
	u2 := NewReaderUser("jer")

	u1.Observer(subject)
	u2.Observer(subject)

	subject.SetState(1)

	subject.SetState(2)
}
