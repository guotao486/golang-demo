/*
 * @Author: GG
 * @Date: 2022-12-08 17:33:33
 * @LastEditTime: 2022-12-08 18:00:50
 * @LastEditors: GG
 * @Description: 组合模式
 * @FilePath: \设计模式\组合模式\main.go
 *
 */
package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
)

/*
 * Composite Pattern 组合模式，又叫部分整体模式。用于把一组相似的对象当做一个单一的对象。

 * 组合模式将对象组合成树形结构，以表示“部分-整体”的层次结构。使得用户对单个对象的使用具有一致性。
 */

// 职员类
type Employee struct {
	Name         string     // 姓名
	Dept         string     // 部门
	Salary       int        // 薪水
	Subordinates *list.List // 属下
}

// 实例化职员类
func NewEmployee(name string, dept string, salary int) *Employee {
	sub := list.New()
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: sub,
	}
}

// 添加职员下属
func (e Employee) Add(emp Employee) {
	// 插入链表的最后一个位置
	e.Subordinates.PushBack(emp)
}

// 删除职员下属
func (e Employee) Remove(emp Employee) {
	// e.Subordinates.Front()  获取链表第一个元素或者nil
	// e.Subordinates.Front().Next() 获取链表的后一个元素或者nil
	// Remove删除链表中的元素e，并返回e.Value。
	// reflect.DeepEqual 判断是否相等
	for i := e.Subordinates.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, emp) {
			e.Subordinates.Remove(i)
		}
	}
}

// 获取职员下属列表
func (e Employee) GetSubordinates() *list.List {
	return e.Subordinates
}

//ToString 获取职员的string信息
func (e *Employee) ToString() string {
	return "[ Name: " + e.Name + ", dept: " + e.Dept + ", Salary: " + strconv.Itoa(e.Salary) + " ]"
}
func main() {

	ceo := NewEmployee("老李", "ceo", 9999)

	pm := NewEmployee("ceo下属张三", "技术", 1000)
	programmer1 := NewEmployee("张三下属李四", "技术", 8000)
	programmer2 := NewEmployee("张三下属王五", "技术", 8000)

	minister := NewEmployee("ceo下属赵六", "财务", 5000)

	finance1 := NewEmployee("赵六下属陈七", "财务", 3000)
	finance2 := NewEmployee("赵六下属牛八", "财务", 2900)

	ceo.Add(*pm)
	ceo.Add(*minister)

	pm.Add(*programmer1)
	pm.Add(*programmer2)

	minister.Add(*finance1)
	minister.Add(*finance2)

	//打印所有职员
	fmt.Println(ceo.ToString())

	fmt.Println("####################")

	for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee) // 类型断言
		fmt.Println(em.ToString())
		fmt.Println("************************")
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
		fmt.Println("----------------")
	}
}
