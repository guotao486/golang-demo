/*
 * @Author: GG
 * @Date: 2022-12-05 16:54:55
 * @LastEditTime: 2022-12-05 17:22:52
 * @LastEditors: GG
 * @Description: 单例模式
 * @FilePath: \设计模式\单例模式\main.go
 *
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Singleton interface {
	dosomething()
}

// 首字母小写 私有的 不能导出
type singleton struct {
	Name string
	Age  int
}

func (s *singleton) dosomething() {
	fmt.Println("do some thing")
}

var (
	once     sync.Once // 保证函数只执行一次
	instance *singleton
)

func GetInstance() Singleton {
	once.Do(
		func() {
			age := rand.Intn(50)
			fmt.Printf("age: %v\n", age)
			instance = &singleton{Name: "tom", Age: age}
		},
	)
	return instance
}

func main() {
	s1 := GetInstance()
	fmt.Printf("s1: %v\n", s1)
	s2 := GetInstance()
	fmt.Printf("s2: %v\n", s2)
}
