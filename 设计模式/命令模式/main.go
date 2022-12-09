/*
 * @Author: GG
 * @Date: 2022-12-09 15:13:41
 * @LastEditTime: 2022-12-09 16:02:26
 * @LastEditors: GG
 * @Description: 命令模式
 * @FilePath: \设计模式\命令模式\main.go
 *
 */
package main

import (
	"container/list"
	"fmt"
)

/*
 * 命令模式(Command Pattern)是一种数据驱动的设计模式。
 * 属于行为型模式。请求以命令的形式包裹在对象中，并传给调用对象。
 * 调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象。该对象执行命令。
 *
 * 任务队列
 */

// 任务接口
type Task interface {
	Execute() bool // 执行方法
}

type Order struct{}

func NewOrder() *Order {
	return &Order{}
}

func (o Order) Execute() bool {
	fmt.Println("订单任务。。。")
	return false
}

type Product struct{}

func NewProduct() *Product {
	return &Product{}
}

func (p Product) Execute() bool {
	fmt.Println("商品任务....")
	return true
}

// 命令调用类
type Broker struct {
	TaskList *list.List
}

// 实例化调用类
func NewBroker() *Broker {
	return &Broker{
		TaskList: list.New(),
	}
}

// 添加任务
func (b *Broker) TakeAdd(take Task) {
	b.TaskList.PushBack(take)
}

func (b *Broker) TakePlace() {
	for i := b.TaskList.Front(); i != nil; {
		// 处理完一个任务就删除掉
		nextTake := i.Next()
		// 正在处理，处理失败应该加入任务尾部
		if !i.Value.(Task).Execute() {
			b.TakeAdd(i.Value.(Task))
		}
		// 删除任务
		b.TaskList.Remove(i)
		i = nextTake
	}
}
func main() {
	broker := NewBroker()

	o := NewOrder()
	broker.TakeAdd(o)

	p := NewProduct()
	broker.TakeAdd(p)

	broker.TakePlace()
}
