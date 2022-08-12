/*
 * @Author: GG
 * @Date: 2022-07-29 16:19:13
 * @LastEditTime: 2022-07-29 16:40:51
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_os_file.go
 *
 */
package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

//读文件
func readFile() {
	s, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("s: %v\n", string(s[:]))
	}
}

// 写入文件，覆盖
func writeFile() {
	s := "我不好1234"
	err := os.WriteFile("test.txt", []byte(s), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 重命名文件
func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func main1() {
	// createFile()
	// readFile()
	// writeFile()
	// renameFile()
}
