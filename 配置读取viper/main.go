/*
 * @Author: GG
 * @Date: 2022-09-16 16:29:37
 * @LastEditTime: 2022-09-16 17:07:59
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

// 单个读取
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

// 将读取的数据转换成结构体对象
type Config struct {
	Db DB `mapstructrue:"db"`
}
type DB struct {
	Username string `mapstructrue:"username"`
	Password string `mapstructrue:"password"`
}

func ReadIniToStruct() {
	config := Config{}
	v := viper.New()
	path, _ := os.Getwd()
	v.AddConfigPath(path + "/配置读取viper/") // 路径
	v.SetConfigName("config")             // 文件名
	v.SetConfigType("ini")                // 文件类型

	err := v.ReadInConfig() // 读取配置
	if err != nil {
		fmt.Println(err)
	}

	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct %w \n", err))
	}
	fmt.Printf("config: %v\n", config)
	fmt.Printf("config.Db: %v\n", config.Db)
}
func main() {

	// ReadIni()
	ReadIniToStruct()
}
