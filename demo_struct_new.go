/*
 * @Author: GG
 * @Date: 2022-07-26 15:56:32
 * @LastEditTime: 2022-07-26 16:06:35
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_struct_new.go
 *
 */
package main

import (
	f "fmt"
) //包别名

type s_person struct {
	name string
	age  int
}

func newPerson(name string, age int) (*s_person, error) {
	if name == "" {
		return nil, f.Errorf("name not nil")
	}

	if age < 0 {
		return nil, f.Errorf("age < 0")
	}
	return &s_person{name: name, age: age}, nil
}

func main1() {
	// 构造方法
	s, e := newPerson("tom", 18) //s: &{tom 18}
	f.Printf("s: %v\n", s)
	f.Printf("e: %v\n", e)

	s1, e1 := newPerson("", -1)
	f.Printf("s1: %v\n", s1)
	f.Printf("e1: %v\n", e1)

}
