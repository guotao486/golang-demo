/*
 * @Author: GG
 * @Date: 2022-07-29 10:14:36
 * @LastEditTime: 2022-07-29 16:19:01
 * @LastEditors: GG
 * @Description: os 目录操作
 * @FilePath: \golang-demo\demo_os_dir.go
 *
 */
package main

import (
	"fmt"
	"os"
)

// 创建目录
func createDir() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("mkdir test success")
}

// 创建多层目录
func createDirAll() {
	err := os.MkdirAll("test/a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("mkdirall test/a/b/c success")
}

// 删除目录
func removeDir() {
	err := os.Remove("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("remove dir test success")
}

// 删除多层目录
func removeDirAll() {
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("removeall dir test success")
}

// 当前目录
func pwd() {
	dir, err := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	fmt.Printf("err: %v\n", err)
}

// 切换当前目录
func cdd() {
	err := os.Chdir("d:/")
	fmt.Printf("err: %v\n", err)
}

// 获得临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}
func main1() {

	// createDir()
	// removeDir()
	// createDirAll()
	// removeDirAll()
	// cdd()
	// pwd()
	getTemp()
}
