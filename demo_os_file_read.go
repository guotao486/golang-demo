/*
 * @Author: GG
 * @Date: 2022-08-01 15:43:42
 * @LastEditTime: 2022-08-01 17:02:18
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_os_file_read.go
 *
 */
package main

import (
	"fmt"
	"os"
)

func openClose() {
	// 读取文件,只读权限
	f, err := os.Open("test2.txt")
	fmt.Printf("f: %v\n", f.Name())
	fmt.Printf("err: %v\n", err)

	err = f.Close()
	fmt.Printf("err: %v\n", err)

	// 读取文件，若文件不存在则创建
	f, err = os.OpenFile("test1.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f.Name(): %v\n", f.Name())
	fmt.Printf("err: %v\n", err)

	err = f.Close()
	fmt.Printf("err: %v\n", err)
}

func readOps() {
	// 循环读取
	// f, _ := os.Open("test2.txt")
	// for {
	//  缓冲区
	// 	buf := make([]byte, 6)
	//  读取并放置缓冲区内，返回长度和错误信息
	// 	n, err := f.Read(buf)
	// 	fmt.Printf("buf: %v\n", string(buf))
	// 	fmt.Println(n)
	//	检查是否读完
	// 	if err == io.EOF {
	// 		break
	// 	}

	// }
	// f.Close()
	// 循环读取 end

	// 定位读取
	// f, _ := os.Open("test2.txt")
	// buf := make([]byte, 10)
	// // 从第五位开始读取10个字节
	// n, err := f.ReadAt(buf, 5)
	// fmt.Printf("n: %v\n", n)
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("buf: %v\n", string(buf))
	// f.Close()
	// 定位读取end

	// 读取目录
	// f, _ := os.Open("a/b")
	// // 大于0，返回参数数量的目录，小于等于0返回所有，返回切片
	// de, err := f.ReadDir(1)

	// fmt.Printf("err: %v\n", err)
	// for _, v := range de {
	// 	fmt.Printf("v.IsDir(): %v\n", v.IsDir())
	// 	fmt.Printf("v.Name(): %v\n", v.Name())
	// }
	// f.Close()
	// 读取目录end

	// 定位
	f, _ := os.Open("test2.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()
}
func main1() {

	// openClose()
	readOps()
}
