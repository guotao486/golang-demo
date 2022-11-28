/*
 * @Author: GG
 * @Date: 2022-08-29 15:29:05
 * @LastEditTime: 2022-08-29 16:26:21
 * @LastEditors: GG
 * @Description:http get
 * @FilePath: \golang-demo\http\get\main.go
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func get_test_1() {
	url := "http://apis.juhe.cn/simpleWeather/query?key=087d7d10f700d20e27bb753cd806e40b&city=北京"
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
}

func get_test_2() {
	// 参数结构体
	params := url.Values{}
	// 解析url
	Url, err := url.Parse("http://apis.juhe.cn/simpleWeather/query")
	if err != nil {
		log.Fatal(err)
	}

	// 添加参数
	params.Set("key", "087d7d10f700d20e27bb753cd806e40b")
	params.Set("city", "北京")
	// 参数转码
	Url.RawQuery = params.Encode()
	// 获取完整的url
	urlPath := Url.String()
	fmt.Printf("urlPath: %v\n", urlPath)
	r, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
	jsonParse(b)
}

func jsonParse(b []byte) {
	type result struct {
		Reason    string                 `json:"reason"`
		Result    map[string]interface{} `json:"result"`
		ErrorCode string                 `json:"error_code"`
	}

	var res result
	_ = json.Unmarshal(b, &res)
	fmt.Printf("res: %v\n", res)
}

func get_test_add_header() {
	client := &http.Client{}
	// 请求对象
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	// 添加header属性
	req.Header.Add("name", "tom")
	req.Header.Add("age", "25")
	// 执行请求
	r, _ := client.Do(req)
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Printf("body: %v\n", string(body))
}
func main() {
	// get_test_1()
	// get_test_2()
	get_test_add_header()
}
