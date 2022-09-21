/*
 * @Author: GG
 * @Date: 2022-09-21 17:18:03
 * @LastEditTime: 2022-09-21 17:29:14
 * @LastEditors: GG
 * @Description:
 * @FilePath: \加密解密bcrypt\main.go
 *
 */
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost) // 加密处理
	return hash, err
}

// pwd1 是加密后的字符串
func ComparePwd(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	} else {
		return true
	}
}
func main() {
	p, e := GenPwd("1234")
	fmt.Printf("p: %v\n", string(p))
	fmt.Printf("e: %v\n", e)
	r := ComparePwd("$2a$04$BEVkeo.XdyaldOVjx7zcNOCfGNOQ.nfFQO46upO/1hJHwTBDGmbEG", "12345")
	fmt.Printf("r: %v\n", r)
}
