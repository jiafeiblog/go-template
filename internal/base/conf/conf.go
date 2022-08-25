package conf

import (
	"time"
)

// GlobalConfig 全局配置索引
var GlobalConfig *Conf

// Conf 配置文件映射
type Conf struct {
	Base   Base
	Logger Logger
	Repo   Repo
	Redis  Redis
}

// Base 基本配置，包括相关端口号地址等
type Base struct {
	WebPort         string // web 服务端口
	PRCPort         string // rpc 服务端口
	RegisterAddress string // 注册中心地址
	ServerAddress   string // 服务地址
	Timezone        string // 时区
}

// Logger 日志配置
type Logger struct {
	Level        string        // 日志打印级别
	Path         string        // 日志存放路径
	MaxAge       time.Duration // 最大存放时间
	RotationTime time.Duration // 日志分割时间
	StdOut       bool          // 控制台输出
}

// Repo 数据仓库配置
type Repo struct {
	Connection      string // Data Source Name
	ConnMaxLifeTime int    // 连接池中每个连接的最大生存时间，单位秒。
	MaxOpenConn     int    // 连接池中允许同时打开的最大连接数
	MaxIdleConn     int    // 连接池中允许存在的最大空闲连接数
}

// Redis 配置
type Redis struct {
	Connection string // Data Source Name
	Username   string
	Password   string
}
