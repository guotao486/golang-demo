/*
 * @Author: GuoTao
 * @Date: 2023-11-08 09:54:44
 * @LastEditTime: 2023-11-08 10:39:44
 * @LastEditors: GuoTao
 * @Description:
 * @FilePath: \缓存淘汰算法\cache\cache.go
 *
 */
package cache

import (
	"fmt"
	"runtime"
)

// cache 缓存接口
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest() // 删除无用的缓存
	Len() int
}

type Value interface {
	Len() int
}

// 计算内存
func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, uint8, int8:
		n = 1
	case uint16, int16:
		n = 2
	case uint32, int32, float32:
		n = 4
	case uint64, int64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.Value", value))
	}

	return n
}
