/*
 * @Author: GG
 * @Date: 2022-07-25 21:49:58
 * @LastEditTime: 2022-07-26 15:02:15
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_struct_func.go
 *
 */
package main

import "fmt"

type f_person struct {
	name string
}

func (p f_person) getName() {
	fmt.Printf("p.name: %v\n", p.name)
}
func (p *f_person) setName(name string) {
	p.name = name
	fmt.Printf("p.name: %v\n", p.name)
}
func main1() {
	var per f_person
	per.name = "tom"
	fmt.Printf("per: %v\n", per)
	per.getName()
	per.setName("bin")

	per1 := f_person{name: "123"}
	per1.setName("321")

	per2 := new(f_person)
	per2.setName("111")
}
