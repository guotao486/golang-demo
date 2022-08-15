/*
 * @Author: GG
 * @Date: 2022-08-15 14:20:04
 * @LastEditTime: 2022-08-15 16:28:10
 * @LastEditors: GG
 * @Description: xml 编码解码
 * @FilePath: \golang-demo\encoding_xml\demo_xml.go
 *
 */
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	sex     int      `xml:"age"`
}

// 将结构体转换成xml
func marshal() {
	p := person{
		Name: "tom",
		Age:  18,
		sex:  1,
	}

	r, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Printf("%v", string(r))
}

// 将xml转换成结构体
func unmarshal() {
	x := `
<person>
	<name>tom</name>
	<age>18</age>
  </person>
	`
	b := []byte(x)

	var p person
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

// 写入文件
func write() {
	// 要写入的结构体原文
	p := person{
		Name: "tom",
		Age:  18,
		sex:  1,
	}
	// 打开文件句柄
	f, _ := os.OpenFile("./a.xml", os.O_RDONLY|os.O_CREATE, 0777)
	// 执行完成关闭文件
	defer f.Close()
	// 传入文件
	x := xml.NewEncoder(f)
	// 开始编码并写入
	x.Encode(p)
}

// 读取文件
func read() {
	// 读取文件内容
	f, _ := os.ReadFile("a.xml")
	// 要转换的格式
	var p person
	// 开始解码
	xml.Unmarshal(f, &p)
	fmt.Printf("p: %v\n", p)
}
func main() {
	// marshal()
	// unmarshal()
	// write()
	read()
}
