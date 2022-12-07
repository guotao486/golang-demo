/*
 * @Author: GG
 * @Date: 2022-12-07 16:11:35
 * @LastEditTime: 2022-12-07 16:26:59
 * @LastEditors: GG
 * @Description: 建造者模式
 * @FilePath: \设计模式\建造者模式\main.go
 *
 */
package main

import "fmt"

// 主要解决在软件系统中，有时候面临一个复杂对象的创建工作，通常这个复杂对象由各个部分的子对象用一定的算法构建成。
// 由于需求的变化，这个复杂对象的各个部分通常会出现巨大的变化，所以，将各个子对象独立出来，容易修改

// 接口
type Builder interface {
	buildDisk()
	buildCPU()
	buildRom()
}

// 子对象1
type SuperComputer struct {
	Name string
}

func (b SuperComputer) buildDisk() {
	fmt.Println("超大硬盘")
}

func (b SuperComputer) buildCPU() {
	fmt.Println("超快CPU")
}

func (b SuperComputer) buildRom() {
	fmt.Println("超大内存")
}

// -----end-------

// 子对象2
type LowComputer struct {
	Name string
}

func (b LowComputer) buildDisk() {
	fmt.Println("超小硬盘")
}

func (b LowComputer) buildCPU() {
	fmt.Println("超慢CPU")
}

func (b LowComputer) buildRom() {
	fmt.Println("超小内存")
}

// --------end--------

// 调用类
type Drictor struct {
	builder Builder
}

// 实例化调用类
func NewConstruct(b Builder) *Drictor {
	return &Drictor{
		builder: b,
	}
}

// 构造
func (d Drictor) Construct() {
	d.builder.buildCPU()
	d.builder.buildDisk()
	d.builder.buildRom()
}

func main() {
	// 构造子对象1
	sc := SuperComputer{}
	d := NewConstruct(&sc)
	d.Construct()

	fmt.Println("----------")
	// 构造子对象2
	lc := LowComputer{}
	d2 := NewConstruct(&lc)
	d2.Construct()
}
