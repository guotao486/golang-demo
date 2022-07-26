/*
 * @Author: GG
 * @Date: 2022-07-26 20:49:03
 * @LastEditTime: 2022-07-26 20:54:06
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goroutines_mutex.go
 *
 */
package main

import (
	"fmt"
	"sync"
)

var m int = 100
var wp sync.WaitGroup
var lock sync.Mutex

func main1() {

	add := func() {
		defer wp.Done()
		lock.Lock() //加锁
		m += 1
		lock.Unlock() // 释放锁
	}
	sub := func() {
		defer wp.Done()
		lock.Lock()
		m -= 1
		lock.Unlock()
	}

	for i := 0; i < 100; i++ {
		wp.Add(1)
		go add()
		wp.Add(1)
		go sub()
	}

	wp.Wait()
	// 因为协程任务系统随机分配的原因会导致公共变量最后结果与预期不符合
	// 对变量加锁解决并发问题
	fmt.Printf("m: %v\n", m)
}
