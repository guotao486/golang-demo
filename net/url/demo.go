/*
 * @Author: GG
 * @Date: 2023-04-28 10:00:41
 * @LastEditTime: 2023-04-28 10:15:02
 * @LastEditors: GG
 * @Description:
 * @FilePath: \url\demo\demo1.go
 *
 */
package main

import (
	"fmt"
	"net/url"
)

func main() {
	test4()
}

// 解析url
func test1() {
	u, err := url.Parse("http://www.example.com/search?q=golang#top")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("http://www.example.com/search?q=golang#top")
	fmt.Printf("u.Scheme: %v\n", u.Scheme)
	fmt.Printf("u.Host: %v\n", u.Host)
	fmt.Printf("u.Path: %v\n", u.Path)
	fmt.Printf("u.Query(): %v\n", u.Query())
	fmt.Printf("u.Fragment: %v\n", u.Fragment)
}

// 构建
func test2() {
	u := &url.URL{
		Scheme:   "https",
		Host:     "www.baidu.com",
		Path:     "search",
		Fragment: "top",
	}

	q := u.Query()
	q.Set("q", "golang")
	q.Set("p", "1")
	u.RawQuery = q.Encode()

	fmt.Printf("u.String(): %v\n", u.String())
}

// 解析查询参数
func test3() {
	values, err := url.ParseQuery("q=golang&page=1&limit=20")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	for k, v := range values {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

	fmt.Printf("values.Get(\"q\"): %v\n", values.Get("q"))
	fmt.Printf("values.Get(\"page\"): %v\n", values.Get("page"))
	fmt.Printf("values.Get(\"limit\"): %v\n", values.Get("limit"))
}

// url 编码解码
func test4() {
	// 编码字符串
	// encoded: http%3A%2F%2Fwww.example.com%2Fsearch%3Fq%3Dgolang%23top

	encoded := url.QueryEscape("http://www.example.com/search?q=golang#top")
	fmt.Printf("encoded: %v\n", encoded)

	// 解码字符串
	// uncoded: http://www.example.com/search?q=golang#top
	uncoded, err := url.QueryUnescape(encoded)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("uncoded: %v\n", uncoded)

}
