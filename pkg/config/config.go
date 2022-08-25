package config

import (
	"github.com/spf13/viper"
)

// InitConfig 读取配置文件，读取配置文件异常直接panic提示 path: 文件路径 c: 配置项结构体&GlobalConfig
func InitConfig(path string, c interface{}) {
	configVip := viper.New()
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	// 配置映射到结构体
	if err := configVip.Unmarshal(c); err != nil {
		panic(err)
	}
}
