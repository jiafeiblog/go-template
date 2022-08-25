package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"template/internal/base/conf"
	"template/internal/base/router"
	"template/pkg/cache"
	"template/pkg/config"
	"template/pkg/db"
	"template/pkg/logger"
	"time"
)

const (
	banner            = "alcohol-all v%v build %v\n"
	defaultConfigPath = "/etc/alcohol-all.yaml"
)

var (
	version   = "dev"
	buildTime = "unknown"

	printVersion = flag.Bool("v", false, "Print version and exit")
	configPath   = flag.String("c", defaultConfigPath, "config file path")
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(banner, version, buildTime))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *printVersion {
		flag.Usage()
		os.Exit(0)
	}

	confInit()
	timezoneInit()
	logInit()
	dbInit()
	//rpcInit()
	webInit()
}

func timezoneInit() {
	if len(conf.GlobalConfig.Base.Timezone) > 0 {
		time.Local, _ = time.LoadLocation(conf.GlobalConfig.Base.Timezone)
	}
}

func logInit() {
	logConf := conf.GlobalConfig.Logger
	maxAge := 7 * 24 * time.Hour
	rotationTime := 24 * time.Hour
	logger.InitLogger("shopping-goods", logConf.Level, logConf.Path, maxAge, rotationTime, logConf.StdOut)
}

func confInit() {
	config.InitConfig(*configPath, &conf.GlobalConfig)
}

func dbInit() {
	repoConf := conf.GlobalConfig.Repo
	db.InitDB(repoConf.Connection)
	// 入库均为 UTC 时间
	db.Engine.DatabaseTZ = time.UTC
	db.Engine.TZLocation = time.UTC

	redisConf := conf.GlobalConfig.Redis
	cache.InitRedis(redisConf.Connection, redisConf.Username, redisConf.Password)
}

func webInit() {
	baseConf := conf.GlobalConfig.Base
	router.InitWebRouter(baseConf.WebPort)
}

//func rpcInit() {
//	baseConf := conf.GlobalConfig.Base
//	router.InitRPCRouter(baseConf.PRCPort)
//	discovery.RegisterService(baseConf.RegisterAddress, discovery.ShoppingGoodsServiceServerName, baseConf.ServerAddress, baseConf.PRCPort)
//	server.InitConsulRPCConn(baseConf.RegisterAddress)
//}
