package main

import (
	routes "github.com/CreFire/rain/api/routers"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/utils/config"
	"github.com/CreFire/rain/utils/log"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置文件读取
	config.ReadConfig()  // 加载配置
	SetConfig()          // 设置初始配置
	log.InitLog()        // 日志初始化
	dal.NewDB()          // 数据库初始化
	routes.SetupRouter() // 设置路由
}

func SetConfig() {
	// 设置Gin的模式
	gin.SetMode(gin.ReleaseMode)
}
