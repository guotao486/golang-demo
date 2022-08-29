/*
 * @Author: GG
 * @Date: 2022-08-29 16:27:04
 * @LastEditTime: 2022-08-29 17:23:06
 * @LastEditors: GG
 * @Description:http post
 * @FilePath: \golang-demo\http\post\main.go
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
	"strings"
	"time"
)

func post_test_1() {
	path := "http://apis.juhe.cn/simpleWeather/query"
	postData := url.Values{}
	postData.Add("key", "087d7d10f700d20e27bb753cd806e40b")
	postData.Add("city", "武汉")
	r, err := http.PostForm(path, postData)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	res, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("res: %v\n", string(res))
}

func post_test_2() {
	postData := url.Values{
		"Name": {"tom"},
		"Age":  {"188"},
	}

	reqBody := postData.Encode()
	res, err := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	r, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("r: %v\n", string(r))
}

// json
func post_test_3() {
	data := make(map[string]interface{})
	data["name"] = "tom"
	data["age"] = "18"
	byteData, _ := json.Marshal(data)
	resp, err := http.Post("http://httpbin.org/post", "application/json", strings.NewReader(string(byteData)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("r: %v\n", string(r))
}

// 自定义
func post_test_4() {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	url := "http://apis.juhe.cn/simpleWeather/query?key=087d7d10f700d20e27bb753cd806e40b&city=北京"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("referer", "http://apis.juhe.cn/")
	res, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("b: %v\n", string(b))

}
func main() {
	// post_test_1()
	// post_test_2()
	post_test_3()
}
