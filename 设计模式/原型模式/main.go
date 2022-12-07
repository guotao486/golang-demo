/*
 * @Author: GG
 * @Date: 2022-12-07 17:06:22
 * @LastEditTime: 2022-12-07 17:21:40
 * @LastEditors: GG
 * @Description: 原型模式
 * @FilePath: \设计模式\原型模式\main.go
 *
 */
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 原型模式用于创建重复的对象。当一个类在创建时开销比较大时(比如大量数据准备，数据库连接)，我们可以缓存该对象，当下一次调用时，返回该对象的克隆。
// 用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。通过实现克隆clone()操作，快速的生成和原型对象一样的实例

type CPU struct {
	Name string
}

type ROM struct {
	Name string
}

type Disk struct {
	Name string
}

type Computer struct {
	Cpu  CPU
	Disk Disk
	Rom  ROM
}

func (c Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, c); err != nil {
		panic(err.Error())
	}
	return pc
}

// 利用gob二进制编码解码实现深拷贝
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	// 编码
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	// 解码
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main() {
	cpu := CPU{"奔腾586"}
	rom := ROM{"金士顿"}
	disk := Disk{"三星"}

	c := &Computer{
		Cpu:  cpu,
		Rom:  rom,
		Disk: disk,
	}

	c1 := c.BackUp()
	fmt.Printf("c1: %v\n", *c1)
	fmt.Printf("c: %v\n", &c)
	fmt.Printf("c1: %v\n", &c1)
}
