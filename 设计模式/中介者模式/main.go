/*
 * @Author: GG
 * @Date: 2022-12-09 17:22:26
 * @LastEditTime: 2022-12-09 17:33:35
 * @LastEditors: GG
 * @Description: 中介者模式
 * @FilePath: \设计模式\中介者模式\main.go
 *
 */
package main

import (
	"fmt"
	"time"
)

/*
 * 中介者模式(Mediator Pattern)用来降低多个对象和类之间的通信复杂性。
 * 对象之间存在大量的关联关系，势必会导致系统结构变得复杂。若一个对象改变，还要跟踪与其相关的对象，跟着一起改变。
 * 中介者模式提供了一个中介类，该类通常处理不同类之间的通信，并且支持松耦合，使代码易于维护。
 * 类似实例，各个国家有事情，找联合国调停，彼此之间不互相接触
 */

// 联合国
type UnitedNations struct{}

var un = NewUniteNations()

func NewUniteNations() *UnitedNations {
	return &UnitedNations{}
}

func (un *UnitedNations) SendMessage(nation *Nation, msg string) {
	fmt.Printf("%s: [ %s ]: %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		nation.Name,
		msg)
}

// 其它国家
type Nation struct {
	Name string
}

func NewNation(name string) *Nation {
	return &Nation{
		Name: name,
	}
}

func (n *Nation) SendMessage(msg string) {
	un.SendMessage(n, msg)
}

func main() {
	n1 := NewNation("中国")
	n1.SendMessage("我要强大...")

	n2 := NewNation("美国")
	n2.SendMessage("我要看看谁不服...")
}
