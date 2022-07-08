package helper

import (
	"GoEngine/entity"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//定义配置变量
var config *entity.Config

//定义配置文件路径变量
var configFilePath string

//实例化接收命令行参数
func init() {
	//设置配置文件路径
	flag.StringVar(&configFilePath, "config", "./config/engine-dev.yml", "config file path")

	//解析参数
	flag.Parse()
}

//获取配置信息对象
func GetConfig() *entity.Config {
	//初始化配置信息
	if config == nil {
		//定义配置变量
		config = &entity.Config{}

		//读取配置文件内容
		data, _ := ioutil.ReadFile(configFilePath)

		//解析配置信息
		yaml.Unmarshal([]byte(data), &config)
	}

	return config
}
