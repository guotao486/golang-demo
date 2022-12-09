/*
 * @Author: GG
 * @Date: 2022-12-09 09:58:37
 * @LastEditTime: 2022-12-09 10:14:00
 * @LastEditors: GG
 * @Description: 外观模式
 * @FilePath: \设计模式\外观模式\main.go
 *
 */
package main

import "fmt"

/*
 * 外观模式(Facade Pattern)隐藏系统的复杂性，并向客户端提供了一个客户端可以访问的接口。
 * 外观模式为子系统中的一组接口提供一个一致的界面，这个接口使得这一子系统更加容易使用。
 * 封装的实现
 */

// 接口
type Sport interface {
	Run()
}

type Basketball struct{}

type Football struct{}

// 实例化Basketball
func NewBasketball() *Basketball {
	return &Basketball{}
}

// 实例化Football
func NewFootball() *Football {
	return &Football{}
}

func (b Basketball) Run() {
	fmt.Println("打篮球")
}

func (b Football) Run() {
	fmt.Println("踢足球")
}

// 外观
type SportFacade struct {
	basketball Basketball
	football   Football
}

// 实例化外观
func NewSportFacade() *SportFacade {
	return &SportFacade{
		basketball: Basketball{},
		football:   Football{},
	}
}

// 打篮球，通过外观去调用逻辑
func (s SportFacade) PlayBasketball() {
	s.basketball.Run()
}

// 踢足球
func (s SportFacade) PlayFootball() {
	s.football.Run()
}

func main() {
	sf := NewSportFacade()
	sf.PlayBasketball()
	sf.PlayFootball()
}
