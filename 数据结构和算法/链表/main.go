/*
 * @Author: GG
 * @Date: 2023-01-08 14:58:41
 * @LastEditTime: 2023-01-08 17:09:02
 * @LastEditors: GG
 * @Description: 链表
 * @FilePath: \数据结构和算法\链表\main.go
 *
 */
package main

import "fmt"

type Object interface{}

type Node struct {
	Data Object //定义数据域
	Next *Node  //定义地址域（指向下一个表的地址）
}

type List struct {
	headNode *Node //头节点
}

// 判断列表是否为空，如果头节点为空，则列表为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

// 链表长度
func (this *List) GetLength() int {
	current := this.headNode
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	return len
}

// 返回最后一个链表元素
func (this *List) Last() *Node {
	current := this.headNode
	for current.Next != nil {
		current = current.Next
	}
	return current
}

// 从链表头部添加元素
func (this *List) Add(d Object) *Node {
	node := &Node{Data: d}
	node.Next = this.headNode
	this.headNode = node
	return node
}

// 从链表尾部添加元素
func (this *List) Append(d Object) {
	node := &Node{Data: d}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		current := this.Last()
		current.Next = node
	}
}

// 在链表指定位置插入元素
func (this *List) Insert(index int, d Object) {
	if index < 0 {
		this.Add(d)
	} else if index > this.GetLength() {
		this.Append(d)
	} else {
		pre := this.headNode
		len := 0
		for len < (index - 1) {
			pre = pre.Next
			len++
		}
		node := &Node{Data: d}
		node.Next = pre.Next
		pre.Next = node
	}
}

// 删除链表指定值元素
func (this *List) Remove(d Object) {
	pre := this.headNode

	if pre.Data == d {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == d {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// 删除链表指定位置元素
func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode
	if index <= 0 {
		this.headNode = pre.Next
	} else if index > this.GetLength() {
		fmt.Println("panic: 没有该位置的元素")
		return
	} else {
		// 找到要删除元素的前一位，将Next 修改为 下下一位，就相当于删除
		len := 0
		for len != (index-1) && pre.Next != nil {
			len++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// 查询链表是否包含某个元素
func (this *List) Contains(d Object) bool {
	current := this.headNode
	for current.Next != nil {
		if current.Data == d {
			return true
		}
		current = current.Next
	}
	return false
}

// 遍历所有元素
func (this *List) LoopList() {
	if !this.IsEmpty() {
		current := this.headNode
		for {
			fmt.Printf("current.Data: %v\n", current.Data)
			if current.Next != nil {
				current = current.Next
			} else {
				break
			}
		}
	}
}
func main() {
	list := List{}
	list.Append("Java")
	list.Append("Python")
	list.Append("Golang")
	list.Insert(1, "PHP")

	fmt.Printf("list.GetLenth(): %v\n", list.GetLength())
	fmt.Println("-----------")
	list.LoopList()

	list.Add("before")
	fmt.Println("---------")

	list.LoopList()

	list.Remove("Java")
	list.RemoveAtIndex(1)

	fmt.Println("---------")
	list.LoopList()

	list.RemoveAtIndex(0)
	fmt.Println("----------")
	list.LoopList()
}
