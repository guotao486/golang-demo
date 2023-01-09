/*
 * @Author: GG
 * @Date: 2023-01-09 10:14:21
 * @LastEditTime: 2023-01-09 11:09:18
 * @LastEditors: GG
 * @Description: container/list标准库是双向链表实现
 * @FilePath: \数据结构和算法\container\list标准库\main.go
 *
 */
package main

import (
	"container/list"
	"fmt"
)

// func (l *List) MoveBefore(e, mark *Element) // 向前移动
// func (l *List) MoveAfter(e, mark *Element)  // 先后移动

// func (l *List) MoveToFront(e *Element)  // 移到头部
// func (l *List) MoveToBack(e *Element)   // 移到尾部

// func (l *List) Front() *Element // 获取最前端元素
// func (l *List) Back() *Element // 获取最后端元素

// func (l *List) InsertBefore(v interface{}, mark *Element) *Element // 指定元素前插入元素
// func (l *List) InsertAfter(v interface{}, mark *Element) *Element // 指定元素后插入元素

// func (l *List) PushFront(v interface{}) *Element // 在最前端插入元素
// func (l *List) PushBack(v interface{}) *Element // 在最后端插入元素
// func (l *List) Remove(e *Element) // 删除元素
func main() {
	linkedlist := list.New()

	// 尾部添加
	linkedlist.PushBack("PHP")
	linkedlist.PushBack("java")

	// 头部添加
	linkedlist.PushFront("golang")

	loop(linkedlist)
	fmt.Println("-------")

	remove("java", linkedlist)
	loop(linkedlist)
	fmt.Println("-------")
	linkedlist.InsertAfter("c#", linkedlist.Front())
	loop(linkedlist)
	fmt.Println("-------")
	linkedlist.MoveToBack(linkedlist.Front())
	loop(linkedlist)
}

func loop(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("e.Value: %v\n", e.Value)
	}
}

func remove(el string, list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value == el {
			list.Remove(e)
			break
		}
	}
}
