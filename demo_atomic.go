/*
 * @Author: GG
 * @Date: 2022-07-26 21:49:17
 * @LastEditTime: 2022-07-26 21:58:23
 * @LastEditors: GG
 * @Description:atomic 原子操作,可替代mutex，并更好
 * @FilePath: \golang-demo\demo_atomic.go
 *
 */
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func main1() {
	add := func() {
		atomic.AddInt32(&i, 1)
	}

	sub := func() {
		atomic.AddInt32(&i, -1)
	}

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("i: %v\n", i)

	// 增减操作
	// - func AddInt32(addr *int32, delta int32) (new int32)
	// - func AddInt64(addr *int64, delta int64) (new int64)
	// - func AddUint32(addr *uint32, delta uint32) (new uint32)
	// - func AddUint64(addr *uint64, delta uint64) (new uint64)
	// - func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)

	// 载入操作
	// LoadInt32(addr *int32)(val int32)

	//比较并交换
	// CompareAndSwapInt32(addr *int32, old, new int32)(swapped bool)

	// 直接交换
	// SwapInt32(addr *int32, new int32)(old int32)

	// 存储
	// StoreInt32(addr *int32, val int32)
}
