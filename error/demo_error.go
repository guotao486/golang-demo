/*
 * @Author: GG
 * @Date: 2022-08-12 18:39:57
 * @LastEditTime: 2022-08-12 19:05:19
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\error\demo_error.go
 *
 */
package main

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) error {
	if s == "" {
		return errors.New("字符串不能为空")
	} else {
		return nil
	}
}

// 自定义错误
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Now(),
		"the file system has gone away",
	}
}

func main() {
	// check("hello")
	// err := check("")
	// fmt.Printf("err: %v\n", err.Error())
	err := oops()
	t, e := err.(error)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("e: %v\n", e)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Println(err)
	}

	me := MyError{time.Now(), "12321312"}
	fmt.Printf("me: %v\n", me)

}
