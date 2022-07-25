/*
 * @Author: GG
 * @Date: 2022-07-25 21:53:45
 * @LastEditTime: 2022-07-25 21:54:54
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_func_receiver.go
 *
 */
package main

import "fmt"

//方法的receiver type并非一定要是struct类型,type定义的类型别名、slice、map、channel、func类型等都可以

// struct
type r_person struct {
	name string
}

func (p r_person) getName() {
	fmt.Printf("p.name: %v\n", p.name)
}
func (p *r_person) setName(name string) {
	p.name = name
	fmt.Printf("p.name: %v\n", p.name)
}

// struct end
func main1() {
	var per f_person
	per.name = "tom"
	per.getName()
	per.setName("bin")
}
