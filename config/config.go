package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type MainConfig struct {
	Version string   `json:"version"`
	Url     []string `json:"url"`
}

var Conf *MainConfig

//修改默认配置文件位置的函数
func DefConfig(configfile string) string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, "config.json")
}

//从配置文件中载入json字符串
func LoadConfig(path string) *MainConfig {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	return mainConfig
}

//初始化 可以运行多次
func SetConfig(path string) {
	mainConfig := LoadConfig(path)
	Conf = mainConfig
}
