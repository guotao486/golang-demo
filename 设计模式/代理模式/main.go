/*
 * @Author: GG
 * @Date: 2022-12-09 10:57:17
 * @LastEditTime: 2022-12-09 12:03:24
 * @LastEditors: GG
 * @Description: 代理模式
 * @FilePath: \设计模式\代理模式\main.go
 *
 */
package main

import "fmt"

/*
 * 代理模式(Proxy Pattern)中，一个类代表另一个类的功能。
 * 代理模式为其他对象提供一种代理，以控制对这个对象的访问。
 */

// 接口
type Subject interface {
	Do()
}

type RealSubject struct{}

func (rs RealSubject) Do() {
	fmt.Println("do something......")
}

type Proxy struct {
	Subject
}

func (p Proxy) Do() {
	fmt.Println("代理之前要做的事情")
	p.Subject.Do()
	fmt.Println("代理之后要做的事情")
}

func main() {
	real := RealSubject{}
	proxy := Proxy{real}
	proxy.Do()
}
