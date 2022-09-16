/*
 * @Author: GG
 * @Date: 2022-09-16 10:56:10
 * @LastEditTime: 2022-09-16 13:09:11
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\jwt\main.go
 *
 */
package main

import (
	"fmt"
)

func main1() {
	s, err := GenToken("ghz", "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("s: %v\n", s)

	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjQ4NzAwNTUyLCJpc3MiOiJsYW9ndW8ifQ.eD4c_s5tminPbKJgmCr3n9jUnp0LT2I4t0_Fd5gml7U"

	// mc, err := ParseToken(token)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("mc.Password: %v\n", mc.Password)
	// fmt.Printf("mc.Username: %v\n", mc.Username)
}
