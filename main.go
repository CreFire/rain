package main

import (
	routes "github.com/CreFire/rain/api/routers"
	"github.com/CreFire/rain/utils/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置文件读取
	config.ReadConfig()
	// 设置Gin的模式
	gin.SetMode(gin.DebugMode)
	// 设置路由
	router := routes.SetupRouter()
	// 启动HTTP服务器
	if err := router.Run(config.Conf.Server.Port); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
