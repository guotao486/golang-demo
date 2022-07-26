/*
 * @Author: GG
 * @Date: 2022-07-26 17:57:41
 * @LastEditTime: 2022-07-26 18:02:37
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_goroutines_waitgroup.go
 *
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func wg_show(i int) {
	defer wg.Done() // 执行完计数器-1
	fmt.Printf("i: %v\n", i)
}

func main1() {

	for i := 0; i < 5; i++ {
		wg.Add(1) //计数器+1
		go wg_show(i)
	}

	// 同步阻塞，执行完后执行后面代码
	wg.Wait()
	fmt.Println("end...")
}
