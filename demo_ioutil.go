/*
 * @Author: GG
 * @Date: 2022-08-01 22:51:24
 * @LastEditTime: 2022-08-01 23:01:08
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_ioutil.go
 *
 */
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main1() {
	// 读取数据
	r := strings.NewReader("abcdefg")
	f, _ := ioutil.ReadAll(r)
	fmt.Printf("f: %v\n", string(f))

	// 读取目录
	fi, _ := ioutil.ReadDir("a")
	for _, v := range fi {
		fmt.Printf("v: %v\n", v.Name())
	}

	//读取文件
	ff, _ := ioutil.ReadFile("test1.txt")
	fmt.Printf("ff: %v\n", string(ff))

	//写入文件
	ioutil.WriteFile("test1.txt", []byte("hello wolrd"), 0755)

	//TempDir 创建临时目录，返回路径

	//TempFile 在路径中创建临时文件，返回os.File

}
