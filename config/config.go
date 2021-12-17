package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Configs map[string]json.RawMessage

var Confs Configs

//修改默认配置文件位置的函数
func DefConfig(configfile *string) {
	pwd, _ := os.Getwd()
	*configfile = strings.Join([]string{pwd, "/config.json"}, "")
}

//从配置文件中载入json字符串
func LoadConfig(path string) Configs {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	allConfigs := make(Configs, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	return allConfigs
}

//初始化 可以运行多次
func SetConfig(path string) {
	allConfigs := LoadConfig(path)
	Confs = allConfigs
}
