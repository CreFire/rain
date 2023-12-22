package routes

import (
	"github.com/CreFire/rain/api/handlers"
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/middleware"
	"github.com/CreFire/rain/utils/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()
	visitor := router.Group("/", middleware.RoleMiddleware(common.VISTOR))
	visitor.POST("login", handlers.LoginHandler)
	visitor.POST("register", handlers.RegisterHandler)
	visitor.GET("select", handlers.SelectIdHandler)
	// 租户相关路由
	tenant := router.Group("tenant", middleware.RoleMiddleware(common.CUSTOM))
	tenant.GET("/:id", handlers.GetTenant)
	tenant.POST("/pay", handlers.PayRent)

	// 管理员相关路由
	admin := router.Group("admin", middleware.RoleMiddleware(common.ADMIN))
	admin.GET("/overdue", handlers.GetOverduePayments)
	admin.POST("/notify", handlers.SendNotification)
	// 添加登录路由

	// 微信相关路由
	router.POST("/wechat/notify", handlers.WechatNotify)
	// 启动HTTP服务器
	if err := router.Run(config.Conf.Server.Port); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
