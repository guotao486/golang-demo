/*
 * @Author: GG
 * @Date: 2022-07-25 21:11:03
 * @LastEditTime: 2022-07-25 21:21:14
 * @LastEditors: GG
 * @Description:结构体指针
 * @FilePath: \golang-demo\demo_struct_needle.go
 *
 */
package main

import "fmt"

func main1() {
	type Person struct {
		id   int
		name string
	}

	var tom = Person{1, "tom"}
	//指针类型
	var p_person *Person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %T\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
	fmt.Printf("p_person: %v\n", p_person)

	// new 关键字对结构体实例化，是指针类型
	var p_person2 = new(Person)
	fmt.Printf("p_person2: %v\n", *p_person2)
	fmt.Printf("p_person2: %v\n", p_person2)
	fmt.Printf("p_person2: %T\n", p_person2)

	//结构体指针类型访问使用.
	fmt.Printf("p_person2.id: %v\n", p_person2.id)
	fmt.Printf("p_person2.name: %v\n", p_person2.name)
	p_person.id = 2
	p_person2.name = "bin2"
	fmt.Printf("p_person2: %v\n", p_person2)
}
