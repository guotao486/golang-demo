/*
 * @Author: GG
 * @Date: 2022-07-24 15:09:48
 * @LastEditTime: 2022-07-24 16:02:34
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\demo_string.go
 *
 */
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main1() {
	var str1 string = "hello world"
	var html string = `
	<html>
	<head><title>hello golang</title></head>
	</html>
	`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("html: %v\n", html)

	//字符串连接
	//golang 里面的字符串都是不可变的,每次运算都会产生一个新的字符串
	//所以会产生很多临时的无用的字符串，不仅没有用，还会给 gc 带来额外的负担，所以性能比较差
	name := "tom"
	age := "20"
	msg := name + " " + age
	fmt.Printf("msg: %v\n", msg)
	msg = ""
	msg += name
	msg += " "
	msg += age
	fmt.Printf("msg: %v\n", msg)

	//fmt.Sprintf()函数
	//内部使用 []byte 实现，不像直接运算符这种会产生很多临时的字符串
	//但是内部的逻辑比较复杂，有很多额外的判断，还用到了 interface，所以性能也不是很好
	msg = fmt.Sprintf("%s,%s", name, age)
	fmt.Printf("msg: %v\n", msg)

	//strings.Join()
	//join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入
	//在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小
	msg = strings.Join([]string{name, age}, ",")
	fmt.Printf("msg: %v\n", msg)

	//buffer.WriteString()
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString(age)
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	// 字符串常用方法
	/**
	  *
	    len(str)	求长度
		+或fmt.Sprintf	拼接字符串
		strings.Split	分割
		strings.contains	判断是否包含
		strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
		strings.Index(),strings.LastIndex()	子串出现的位置
		strings.Join(a[]string, sep string)	join操作
	*/
	s := "hello world！"
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasSuffix(s, \"world！\"): %v\n", strings.HasSuffix(s, "world！"))
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))
}
