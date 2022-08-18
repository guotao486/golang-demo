/*
 * @Author: GG
 * @Date: 2022-08-12 11:08:26
 * @LastEditTime: 2022-08-18 09:27:33
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\log\demo_log.go
 *
 */
package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}

// panic 抛出异常
func test2() {
	// defer 还是会执行
	defer log.Print("my log end...")
	log.Print("my log")
	log.Panic("log panic")
	// 后面的不会执行
	fmt.Println("fmt end...")

}

// fatal 打印日志，结束程序
func test3() {
	// defer 不会执行
	defer fmt.Println("fatal defer")
	log.Print("fatal log")
	log.Fatal("fatal")
	// 后面的不会执行
	log.Println("FATAL END")
}

// 日志配置
func test4() {
	log.Print("demo...")
	// Lmicroseconds 微秒
	// Llongfile 文件名和行号
	// Lshortfile 最终的文件名和行号，覆盖Llongfile
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("my log")
}

// 日志前缀设置
func test5() {
	log.SetPrefix("error:")
	log.Print("my log")
}

func test6() {
	f, err := os.OpenFile("log/error.log", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	// 传入文件句柄
	log.SetOutput(f)
	log.Print("my log...")
}
func main() {
	test1()
	test2()
	test3()
	// test4()
	// test5()
	// test6()
}
