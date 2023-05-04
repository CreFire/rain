package main

import (
	"github.com/CreFire/rain/api"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/utils/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// 初始化 Zap 日志记录器
	logger := log.GetLog()

	err := dal.NewDB()
	if err != nil {
		logger.Error("db conn err", log.Err(err))
	}
	r := gin.New()
	// 使用 Zap 记录日志
	r.Use(LoggerMiddleware(logger), LoggerReCover(logger))
	api.Router(r)
	err = r.Run(":8080")
	if err != nil {
		logger.Error("serve failed", log.Err(err))
	}
}

func LoggerReCover(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				l.Error("Panic occurred", log.Any("panic", r))
			}
		}()
		c.Next()
	}
}

func LoggerMiddleware(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求前执行某些操作
		l.Info("New request", zap.String("path", c.Request.URL.Path))
		c.Next()
		// 在请求后执行某些操作
	}
}
