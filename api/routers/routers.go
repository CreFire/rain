package routes

import (
	"github.com/CreFire/rain/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 租户相关路由
	router.GET("/tenant/:id", handlers.GetTenant)
	router.POST("/tenant/pay", handlers.PayRent)

	// 管理员相关路由
	router.GET("/admin/overdue", handlers.GetOverduePayments)
	router.POST("/admin/notify", handlers.SendNotification)

	// 微信相关路由
	router.POST("/wechat/notify", handlers.WechatNotify)

	return router
}
