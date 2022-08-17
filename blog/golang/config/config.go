/*
 * @Author: GG
 * @Date: 2022-08-17 17:19:41
 * @LastEditTime: 2022-08-17 19:43:54
 * @LastEditors: GG
 * @Description:config struct
 * @FilePath: \golang-demo\blog\golang\config\config.go
 *
 */
package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

var Cfg *TomlConfig

type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

func init() {
	Cfg = new(TomlConfig)
	var err error
	Cfg.System.CurrentDir, err = os.Getwd()

	if err != nil {
		panic(err)
	}

	Cfg.System.AppName = "blog golang"
	Cfg.System.Version = 1.0
	_, err = toml.DecodeFile("blog/golang/config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
