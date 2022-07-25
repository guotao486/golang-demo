/*
 * @Author: GG
 * @Date: 2022-07-25 21:03:35
 * @LastEditTime: 2022-07-25 21:50:03
 * @LastEditors: GG
 * @Description:struct 结构体
 * @FilePath: \golang-demo\demo_struct.go
 *
 */
package main

import "fmt"

type person struct {
	id    int
	name  string
	age   int
	email string
	dog   //可以dog1 dog， dog1 为别名
}

type dog struct {
	name string
}

func showPerson(p person) {
	p.id = 1
	p.name = "showPerson"
	fmt.Printf("p: %v\n", p)
}
func showPerson2(p *person) {
	p.id = 3
	p.name = "meiko"
	fmt.Printf("p: %v\n", p)
}
func main1() {
	// 初始化，都是0值
	var tom person
	fmt.Printf("tom: %v\n", tom)
	// 使用
	tom.id = 1
	tom.name = "Tom"
	tom.age = 20
	tom.email = "tom@email.com"
	fmt.Printf("tom: %v\n", tom)

	kite := person{}
	fmt.Printf("kite: %v\n", kite)
	kite.name = "kite"
	fmt.Printf("kite: %v\n", kite)

	ale := person{
		name: "ale",
	}
	fmt.Printf("ale: %v\n", ale)

	// 按字段顺序赋值初始化
	bin := person{
		2,
		"bin",
		20,
		"bin@emial.com",
		dog{
			name: "dog"},
	}
	fmt.Printf("bin: %v\n", bin)

	// 结构体当函数参数, 函数内部修改值不影响外部
	showPerson(bin)
	fmt.Printf("bin: %v\n", bin)
	// 结构体指针当函数参数，会影响到外部值
	showPerson2(&bin)
	fmt.Printf("bin: %v\n", bin)

	bin.dog.name = "dog"
	fmt.Printf("bin: %v\n", bin)
}
