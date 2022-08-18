/*
 * @Author: GG
 * @Date: 2022-08-18 17:53:10
 * @LastEditTime: 2022-08-18 17:53:25
 * @LastEditors: GG
 * @Description:md5
 * @FilePath: \golang-demo\blog\golang\utitl\md5.go
 *
 */
package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

//给字符串生成md5
//@params str 需要加密的字符串
//@params salt interface{} 加密的盐
//@return str 返回md5码
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
