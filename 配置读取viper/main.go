/*
 * @Author: GG
 * @Date: 2022-09-16 16:29:37
 * @LastEditTime: 2022-09-16 16:49:38
 * @LastEditors: GG
 * @Description:配置读取 viper库
 * @FilePath: \golang-demo\配置读取viper\main.go
 *
 */
package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ReadIni() {
	v := viper.New()
	path, _ := os.Getwd()
	v.AddConfigPath(path + "/配置读取viper/") // 路径
	v.SetConfigName("config")             // 文件名
	v.SetConfigType("ini")                // 文件类型

	err := v.ReadInConfig() // 读取配置
	if err != nil {
		fmt.Println(err)
	}

	s := v.GetString("db.username")
	fmt.Printf("s: %v\n", s)
}
func main() {

	ReadIni()
}
