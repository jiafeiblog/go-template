package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"template/pkg/logger"
	"template/pkg/middleware"
)

// InitWebRouter 初始化路由 唯一一个会阻塞的初始化动作，请放置在最后
func InitWebRouter(port string) {
	// 判断是否一release模式启动
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 开发阶段 使用gin的Recovery将日志格式化输出到控制台
	r.Use(gin.Recovery())
	// 生产阶段 使用zap输出到固定文件
	r.Use(middleware.RecoveryWithZap())
	// 解决前端跨域问题
	r.Use(middleware.CorsMiddleware())

	r.GET("/healthz", func(c *gin.Context) { c.String(200, "OK") })

	logger.Infof("web service listening on port %s", port)
	if err := r.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		panic(err)
	}
}
