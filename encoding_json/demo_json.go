/*
 * @Author: GG
 * @Date: 2022-08-15 09:57:25
 * @LastEditTime: 2022-08-15 16:19:31
 * @LastEditors: GG
 * @Description:decoding json 实现json的编码和解码
 * @FilePath: \golang-demo\encoding_json\demo_json.go
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// func Marshal(v interface{}) ([]byte, error)
// 将struct 编码成json，可以接收任何类型

// func Unmarshal(data []byte, v interface{}) error
// 将json 编码成struct

type person struct {
	Name string
	Age  int
	sex  int // 没有外部访问权限则无法被编码
}

// 将结构体转换成json
func marshal() {
	p := person{
		Name: "tom",
		Age:  18,
		sex:  1,
	}

	r, _ := json.Marshal(p)
	fmt.Printf("r: %v\n", string(r))
}

// 将json转换成对应结构体
func unmarshal() {
	// 没有访问权限的字段显示默认值，不存在的字段不处理
	s := []byte(`{"name":"tom","age":18,"sex":1,"email":"asdsad@qq.com"}`)
	var p person
	err := json.Unmarshal(s, &p)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("p: %v\n", p) //p: {tom 18 0}

}

// 将嵌套json类型 解码
func test1() {
	s := []byte(`{"name":"tom","age":18,"sex":1,"person":["kite","tom"]}`)
	// var p interface{}
	var p map[string]interface{} // 该方式可for获取键值对
	err := json.Unmarshal(s, &p)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("p: %v\n", p) // map[age:18 name:tom person:[kite tom] sex:1]

	for k, v := range p {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

// 嵌套类型struct 编码成json
func test2() {
	type person1 struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := person1{
		Name:   "tom",
		Age:    18,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	r, _ := json.Marshal(p)
	fmt.Printf("r: %v\n", string(r)) //r: {"Name":"tom","Age":18,"Email":"tom@gmail.com","Parent":["big tom","big kite"]}
}

// io 读文件内容解码
func test3() {
	// 打开文件
	f, _ := os.Open("a.json")
	// 执行完成后关闭文件
	defer f.Close()
	// 将文件句柄传入
	d := json.NewDecoder(f)
	// 转换后的类型
	var v map[string]interface{}
	// 读取并解码
	d.Decode(&v)

	fmt.Printf("v: %v\n", v)
}

// io 将编码后的内容写入文件
func test4() {
	type person1 struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := person1{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	// 打开文件
	f, _ := os.OpenFile("a.json", os.O_WRONLY, 0777)
	// 执行完成后关闭文件
	defer f.Close()
	// 将文件句柄传入
	e := json.NewEncoder(f)
	// 编码并写入
	e.Encode(p)
}

// io流 Reader Writer 可以扩展到http websocket等场景
func test5() {
	// dec := json.NewDecoder(os.Stdin)
	// a.json : {"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["tom", "kite"]}
	// f, err := os.OpenFile("a.json", os.O_RDONLY|os.O_CREATE, 0644)
	f, err := os.Open("a.json")
	fmt.Printf("err: %v\n", err)
	dec := json.NewDecoder(f)         // 读文件内容 解码
	enc := json.NewEncoder(os.Stdout) // 编码
	for {
		var v map[string]interface{}
		// 将json解码
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		fmt.Println("----------------------")
		fmt.Printf("v: %v\n", v)
		fmt.Println("----------------------")

		//
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}

	/*
		输入 {"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["tom", "kite"]}
		输出
		v: map[Age:20 Email:tom@gmail.com Name:tom Parents:[tom kite]]
		{"Age":20,"Email":"tom@gmail.com","Name":"tom","Parents":["tom","kite"]}
	*/
	/*
		输入 {"Name":"tom","Age":20,"Email":"tom@gmail.com"}
		输出
		v: map[Age:20 Email:tom@gmail.com Name:tom]
		{"Age":20,"Email":"tom@gmail.com","Name":"tom"}
	*/
}

func main() {
	// marshal()
	// unmarshal()
	// test1()
	// test2()
	test3()
	// test4()
}
