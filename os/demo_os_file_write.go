/*
 * @Author: GG
 * @Date: 2022-08-01 17:08:36
 * @LastEditTime: 2022-08-01 17:31:49
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_os_file_write.go
 *
 */
package main

import (
	"os"
)

func main1() {
	// 写入文件，默认在头部覆盖与写入字符串长度相等字节
	// f, _ := os.OpenFile("test1.txt", os.O_RDWR, 0755)
	// n, err := f.Write([]byte("123456"))
	// fmt.Printf("n: %v\n", n)
	// fmt.Printf("err: %v\n", err)

	// 追加
	// f, _ := os.OpenFile("test1.txt", os.O_RDWR|os.O_APPEND, 0755)
	// f.Write([]byte("aaabbb"))

	// 覆盖
	// f, _ := os.OpenFile("test1.txt", os.O_RDWR|os.O_TRUNC, 0755)
	// 直接写入string
	// f.WriteString("hello world")

	// 从指定位置写入
	f, _ := os.OpenFile("test1.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("123"), 5)
	f.Close()
}
