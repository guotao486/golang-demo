/*
 * @Author: GG
 * @Date: 2022-12-09 16:14:11
 * @LastEditTime: 2022-12-09 16:25:28
 * @LastEditors: GG
 * @Description: 迭代器模式
 * @FilePath: \设计模式\迭代器模式\main.go
 *
 */
package main

import "fmt"

type Order struct {
	Name string
}

type OrderList struct {
	Orders []Order
}

func NewOrderList() *OrderList {
	return &OrderList{
		Orders: make([]Order, 0),
	}
}

// 添加
func (ol *OrderList) Add(order Order) {
	ol.Orders = append(ol.Orders, order)
}

func (ol *OrderList) GetIterator() func() (Order, bool) {
	index := 0
	return func() (order Order, ok bool) {
		if index >= len(ol.Orders) {
			return
		}
		order, ok = ol.Orders[index], true
		index++
		return
	}
}
func main() {
	ol := NewOrderList()

	o1 := Order{"book"}
	o2 := Order{"phone"}

	ol.Add(o1)
	ol.Add(o2)

	it := ol.GetIterator()

	for {
		order, ok := it()
		if !ok {
			break
		}
		fmt.Printf("order.Name: %v\n", order.Name)
	}
}
