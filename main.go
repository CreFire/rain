package main

import (
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/internal"
	"github.com/CreFire/rain/tools/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// 初始化 Zap 日志记录器
	logger := log.NewDefault()

	dal.NewDB()
	r := gin.New()
	// 使用 Zap 记录日志
	r.Use(LoggerMiddleware(logger), LoggerReCover(logger))
	internal.Router(r)

}

func LoggerReCover(logger *log.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != any(nil) {
				logger.Fatal("msg:%v",log.Any("panic",any(r)))
			}
		}()
	}
}
func LoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求前执行某些操作
		logger.Info("New request", zap.String("path", c.Request.URL.Path))
		c.Next()
		// 在请求后执行某些操作
	}
}
