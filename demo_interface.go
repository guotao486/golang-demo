/*
 * @Author: GG
 * @Date: 2022-07-26 15:02:27
 * @LastEditTime: 2022-07-26 15:56:23
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_interface.go
 *
 */
package main

import "fmt"

type USB interface {
	read()
	write(string)
}

type video interface {
	play()
}

type dowVideo interface {
	USB
	video
}

type Mobile struct {
	name string
}

type Pad struct {
	name string
}

type iphone struct {
}

func (m Mobile) read() {
	fmt.Println("read....")
}

func (m *Mobile) write(name string) {
	m.name = name
	fmt.Println("write....")
}

func (m Mobile) play() {
	fmt.Println("play...")
}
func (p Pad) read() {
	fmt.Println("read....")
}

func (p *Pad) write(name string) {
	p.name = name
	fmt.Println("write....")
}
func (p Pad) play() {
	fmt.Println("play...")
}

func (i iphone) exc(d dowVideo) {
	d.play()
	d.read()
}
func main1() {
	// 实现接口,引用类型
	m := &Mobile{name: "oppo"}
	var i USB
	i = m
	fmt.Printf("m: %v\n", *m)
	i.read()
	i.write("1加")
	fmt.Printf("m: %v\n", *m)

	// 多个结构体实现一个接口
	m = &Mobile{name: "oppo"}
	p := &Pad{name: "iPad"}
	var v video
	v = m
	v.play()
	v = p
	v.play()

	// 实现多个接口
	p = &Pad{name: "iPad"}
	var ip USB
	var vp video
	ip = p
	vp = p
	ip.read()
	ip.write("iPad 5G")
	vp.play()
	fmt.Printf("p: %T\n", p)

	// 实现多个接口，通过接口嵌套实现
	p = &Pad{name: "iPad"}
	var d dowVideo
	d = p
	d.read()
	d.write("1加")
	d.play()

	// ocp设计原则
	fmt.Println("ocp-------------")
	iphone := iphone{}
	m1 := &Mobile{}
	p1 := &Pad{}
	iphone.exc(m1)
	iphone.exc(p1)
}
