/*
 * @Author: GuoTao
 * @Date: 2023-10-07 16:38:21
 * @LastEditTime: 2023-10-07 20:06:05
 * @LastEditors: GuoTao
 * @Description: 双向链表demo
 * @FilePath: \golang-demo\数据结构和算法\双向链表\main.go
 *
 */
package main

import "fmt"

// 节点数据结构
type DNode struct {
	data interface{}
	prev *DNode // 上一个
	next *DNode // 下一个
}

// 链表数据结构
type DList struct {
	size uint64
	head *DNode // 第一个节点
	tail *DNode // 最后一个节点
}

// 链表的初始化
func InitList() (list *DList) {
	list = new(DList)
	list.head = nil
	list.tail = nil
	list.size = 0
	return
}

// 上一个
func (node *DNode) Prev() *DNode {
	return node.prev
}

// 下一个
func (node *DNode) Next() *DNode {
	return node.next
}

// 获取数据
func (node *DNode) Data() interface{} {
	return node.data
}

// 添加新节点
func (list *DList) Append(data interface{}) {
	// 创建节点
	newData := &DNode{data: data}
	if list.size == 0 { // 第一个节点
		list.head = newData
		list.tail = newData
		newData.prev = nil
		newData.next = nil
	} else { // 接到尾部
		// 新节点的指向修改
		newData.prev = list.tail
		newData.next = nil

		// 之前链表的指向修改
		list.tail.next = newData
		// 更新链表的尾部节点
		list.tail = newData
	}

	// 更新链表计数器
	list.size++
}

// 根据指定位置前面插入节点
/**
 * @description:
 * @param {*DNode} ele
 * @param {interface{}} data
 * @return {*}
 */
func (list *DList) InsertBefore(ele *DNode, data interface{}) bool {
	if ele == nil {
		return false
	}

	if list.IsHead(ele) {
		list.head = &DNode{data: data}
		list.head.next = ele
		list.head.prev = ele.prev
	} else {
		// 创建新节点
		newData := &DNode{data: data}
		// 新节点的指向
		// 新节点上一个就是当前位置节点的上一个节点
		// 新节点的下一个就是当前位置节点
		newData.prev = ele.prev
		newData.next = ele

		// 当前节点的指向修改
		ele.prev = newData
		list.size++
	}
	return true
}

// 在指定位置后面插入节点
func (list *DList) InsertAfter(ele *DNode, data interface{}) bool {
	if ele == nil {
		return false
	}

	if list.IsTail(ele) {
		list.Append(data)
	} else {
		// 创建新节点
		newData := &DNode{data: data}
		// 新节点的指向
		// 新节点上一个就是当前位置节点的上一个节点
		// 新节点的下一个就是当前位置节点
		newData.prev = ele
		newData.next = ele.next

		// 当前节点的指向修改
		ele.next = newData

		list.size++
	}
	return true
}

// 删除节点
func (list *DList) Remove(ele *DNode) bool {
	if ele == nil {
		return false
	}

	// 第一个
	if list.IsHead(ele) {
		list.head = ele.next
	} else if list.IsTail(ele) { // 最后一个
		list.tail = ele.prev
	} else { // 中间
		ele.prev.next = ele.next
		ele.next.prev = ele.prev
		ele.prev = nil
		ele.next = nil
	}

	list.size--
	return true
}

// 判断是否是尾部节点
func (list *DList) IsTail(node *DNode) bool {
	return node == list.tail
}

// 返回尾部节点
func (list *DList) GetTail() *DNode {
	return list.tail
}

// 判断是否头部节点
func (list *DList) IsHead(node *DNode) bool {
	return node == list.head
}

// 返回头部节点
func (list *DList) GetHead() *DNode {
	return list.head
}

// 返回链表长度
func (list *DList) GetSize() uint64 {
	return list.size
}

func main() {
	list := &DList{}
	list.Append("golang")
	list.Append("java")
	list.Append("python")
	list.Append("c++")
	list.Append("c")

	// fmt.Printf("list.GetSize(): %v\n", list.GetSize())

	// list.Remove(list.GetHead())
	// list.Remove(list.GetTail())

	// fmt.Printf("list.GetSize(): %v\n", list.GetSize())

	// p := list.GetHead()
	// for p != nil {
	// 	fmt.Printf("%s ", p.data)
	// 	p = p.Next()
	// }
	// fmt.Println("list end")

	fmt.Printf("list.GetSize(): %v\n", list.GetSize())

	list.InsertAfter(list.GetHead(), "php")
	list.InsertAfter(list.GetTail(), "rust")
	list.InsertBefore(list.GetHead(), "c#")
	list.InsertBefore(list.GetTail(), "go")

	p := list.GetHead()
	for p != nil {
		fmt.Printf("%s ", p.data)
		p = p.Next()
	}
	fmt.Printf("list.GetSize(): %v\n", list.GetSize())
}
