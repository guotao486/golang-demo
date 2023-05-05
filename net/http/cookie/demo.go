/*
 * @Author: GG
 * @Date: 2023-04-28 14:19:42
 * @LastEditTime: 2023-04-28 14:28:09
 * @LastEditors: GG
 * @Description:
 * @FilePath: \http\cookie\demo.go
 *
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	test1()
}

func test1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求中的cookie
		cookie, err := r.Cookie("username")
		if err == nil {
			fmt.Fprintf(w, "Hello: %v\n", cookie.Value)
			return
		}

		cookie = &http.Cookie{
			Name:  "username",
			Value: "Tom",
		}

		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Hello New User!")
	})

	http.ListenAndServe(":8080", nil)
}

// 记住我
func test2() {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// 验证用户名和密码
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "admin" && password == "admin123" {
			// 设置记住我Cookie
			cookie := &http.Cookie{
				Name:    "remember_me",
				Value:   "1",
				Expires: time.Now().Add(30 * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)

			fmt.Fprint(w, "Login successful.")
			return
		}

		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// 检查记住我Cookie
		cookie, err := r.Cookie("remember_me")
		if err == nil && cookie.Value == "1" {
			fmt.Fprint(w, "Welcome back!")
			return
		}

		// 显示登录页面
		fmt.Fprint(w, "Please login.")
	})

	http.ListenAndServe(":8080", nil)

}
