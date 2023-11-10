/*
 * @Author: GuoTao
 * @Date: 2023-11-08 10:07:36
 * @LastEditTime: 2023-11-08 14:23:23
 * @LastEditors: GuoTao
 * @Description:
 * @FilePath: \缓存淘汰算法\cache\fifo\fifo.go
 *
 */
package fifo

import (
	"cache/cache"
	"container/list"
)

// FIFO，先进先出，也就是淘汰缓存中最早添加的记录。
// 在 FIFO Cache 设计中，核心原则就是：如果一个数据最先进入缓存，那么也应该最早淘汰掉。这么认为的根据是，最早添加的记录，其不再被使用的可能性比刚添加的可能性大。
// 这种算法的实现非常简单，创建一个队列（一般通过双向链表实现），新增记录添加到队尾，缓存满了，淘汰队首。

// fifo 是一个 FIFO cache。它不是并发安全的。
type fifo struct {

	// 缓存最大容量，单位字节
	// groupcache 使用的是最大存放 entry 数量
	maxBytes int
	// 当一个 entry 从缓存中移除时是调用该回调函数，默认 nil
	// groupcache 中的 key 是任意的可比较类型； value 是 interface{}
	onEvicted func(key string, value interface{})

	// 已使用的字节数，只包括值，key不算
	usedBytes int

	ll    *list.List
	cache map[string]*list.Element
}

type entry struct {
	key   string
	value interface{}
}

// 获取内存长度
func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}

// 实例化
func New(maxBytes int, onEvicted func(string, interface{})) cache.Cache {
	return &fifo{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}

// Set 往 Cache 尾部增加一个元素（如果已经存在，则移到尾部，并修改值）
func (f *fifo) Set(key string, value interface{}) {
	// 如果key存在
	if e, ok := f.cache[key]; ok {
		f.ll.MoveToBack(e)                                                         // 移动到最后
		en := e.Value.(*entry)                                                     // 获取详情
		f.usedBytes = f.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(value) // 计算已使用总字节 = 总字节 - 当前缓存旧值字节 + 缓存新值字节
		en.value = value                                                           // 修改值
		return
	}

	en := &entry{key, value} // 实例化
	e := f.ll.PushBack(en)   // 在列表尾部插入，并返回列表元素
	f.cache[key] = e         // 将列表元素放入缓存集合

	f.usedBytes += en.Len() // 计算已使用字节
	// 检测是否超出总字节限制，如果超出执行回收机制
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest() // 删除最早添加的缓存
	}
}

// Get 从 cache 中获取 key 对应的值，nil 表示 key 不存在
func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}
	return nil
}

// Del 从 cache 中删除key对应的记录
func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
	return
}

// DelOldest 从 cache 中删除`最旧`的记录
func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Front()) // 删除列表第一个元素
}

// 删除元素
func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	// 从链表中删除
	f.ll.Remove(e)
	en := e.Value.(*entry)
	// 恢复删除元素的使用字节
	f.usedBytes -= en.Len()
	// 从缓存集合中删除
	delete(f.cache, en.key)

	// 删除回调
	if f.onEvicted != nil {
		f.onEvicted(en.key, en.value)
	}
}

// Len 返回当前 cache 的记录数
func (f *fifo) Len() int {
	return f.ll.Len()
}
