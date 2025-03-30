package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

// AppConfig 服务端配置（组合全部配置模型）
type AppConfig struct {
	AppName  string `yaml:"app_name"`
	Port     string `yaml:"port"`
	Mode     string `yaml:"mode"`
	DataBase Mysql  `yaml:"data_base"`
}

// Mysql 配置
type Mysql struct {
	Drive    string `yaml:"drive"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pwd      string `yaml:"pwd"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

// InitConfig 初始化服务器配置
func InitConfig() *AppConfig {
	var config *AppConfig
	file, err := os.Open("./src/configs/config.yaml")
	//file, err := os.Open("../configs/config.yaml")
	if err != nil {
		panic(any(err.Error()))
	}
	// 创建一个新的 YAML 解码器
	decoder := yaml.NewDecoder(file)
	// 使用解码器解码 YAML 文件并将解码的数据存储在变量中
	err = decoder.Decode(&config)
	if err != nil {
		panic(any(err.Error()))
	}
	return config
}
