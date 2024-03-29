/*
 * @Author: GuoTao
 * @Date: 2023-11-08 14:36:13
 * @LastEditTime: 2023-11-08 15:04:08
 * @LastEditors: GuoTao
 * @Description:
 * @FilePath: \缓存淘汰算法\cache\fifo\fifo_test.go
 *
 */
package fifo

import (
	"testing"

	"github.com/matryer/is"
)

func TestSetGet(t *testing.T) {
	is := is.New(t)

	cache := New(24, nil)
	cache.DelOldest()
	cache.Set("k1", 1)
	v := cache.Get("k1")
	is.Equal(v, 1)

	cache.Del("k1")
	is.Equal(0, cache.Len())
}

func TestOnEvicted(t *testing.T) {
	is := is.New(t)
	keys := make([]string, 0, 8)
	onEvicted := func(key string, value interface{}) {
		keys = append(keys, key)
	}

	cache := New(16, onEvicted)

	cache.Set("k1", 1)
	cache.Set("k2", 2)
	cache.Get("k1")

	cache.Set("k3", 3)
	cache.Get("k1")

	cache.Set("k4", 4)

	expected := []string{"k1", "k2"}

	is.Equal(keys, expected)
	is.Equal(2, cache.Len())
}
