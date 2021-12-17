package main

import (
	cnf "CheckHttpsSSL/config"
	"flag"
	"fmt"
	"os"
)

func main() {

	//定义变量接收命令行参数，目前只有一个 -c
	var configfile string
	// StringVar 用名称、控制台参数、默认值、使用信息注册一个string类型flag，并将flag的值保存到configfile指向的变量
	flag.StringVar(&configfile, "c", "", "配置文件路径")
	flag.Parse()

	// 如果configfile变量为空即未指定配置文件，则调用DefConfig函数修改configfile的值为默认配置文件路径
	if configfile == "" {
		cnf.DefConfig(&configfile)
	}
	// 判断默认配置文件是否存在，不存在则退出
	_, err := os.Stat(configfile)
	if os.IsNotExist(err) {
		fmt.Println(err)
		fmt.Println("-c 为空，未指定配置文件")
		os.Exit(2)
	}

	//读取配置文件信息
	cnf.SetConfig(configfile)
	urllist := cnf.Confs["url"]
	fmt.Println(string(urllist))
}
