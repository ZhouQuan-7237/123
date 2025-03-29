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
