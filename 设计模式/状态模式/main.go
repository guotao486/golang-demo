/*
 * @Author: GG
 * @Date: 2022-12-19 11:12:35
 * @LastEditTime: 2022-12-19 11:38:44
 * @LastEditors: GG
 * @Description: 状态模式
 * @FilePath: \设计模式\状态模式\main.go
 *
 */
package main

import "fmt"

// 状态接口
type State interface {
	DoAction(context *Context)
	ToString() string
}

// 状态保存类
type Context struct {
	state State
}

// 实例化 状态保存类
func NewContext() *Context {
	return &Context{
		state: nil,
	}
}

// 设置状态保存类当前的状态
func (c *Context) SetState(s State) {
	c.state = s
}

// 读取状态保存类当前的状态
func (c *Context) GetState() State {
	return c.state
}

// 开始状态类
type StartState struct{}

// 实例化开始状态类
func NewStartState() *StartState {
	return &StartState{}
}

// 开始状态类的DoAction，实现state接口
func (ss *StartState) DoAction(context *Context) {
	fmt.Println("现在是开始状态")
	context.state = ss
}

// 返回开始状态类名称
func (ss *StartState) ToString() string {
	return "开始状态"
}

// 停止状态类
type StopState struct{}

// 实例化停止状态类
func NewStopState() *StopState {
	return &StopState{}
}

// 停止状态类 DoAtion，实现state接口
func (ss *StopState) DoAction(context *Context) {
	fmt.Println("现在是停止状态")
	context.state = ss
}

// 返回停止状态名称
func (ss *StopState) ToString() string {
	return "停止状态"
}
func main() {
	context := NewContext()

	startState := NewStartState()
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())
	fmt.Println("----------------------")
	stopState := NewStopState()
	stopState.DoAction(context)
	fmt.Println(context.GetState().ToString())
}
