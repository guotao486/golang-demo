/*
 * @Author: GG
 * @Date: 2023-04-28 11:25:57
 * @LastEditTime: 2023-04-28 11:26:54
 * @LastEditors: GG
 * @Description:
 * @FilePath: \http\header\cache.go
 *
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Cache-Control", "max-age=60")                                        // 设置缓存时间为60秒
		w.Header().Set("Expires", time.Now().Add(time.Minute).UTC().Format(http.TimeFormat)) // 设置过期时间为1分钟后
		w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))            // 设置最后修改时间为当前时间

		// 如果客户端发送了If-Modified-Since字段，则验证缓存是否过期
		if ifModifiedSince := r.Header.Get("If-Modified-Since"); ifModifiedSince != "" {
			lastModified, err := time.Parse(http.TimeFormat, ifModifiedSince)
			if err == nil && lastModified.UTC().Add(time.Minute).After(time.Now().UTC()) {
				w.WriteHeader(http.StatusNotModified) // 返回304状态码表示内容未修改
				return
			}
		}

		// 返回内容
		fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":8080", nil)
}
