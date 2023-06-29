package main

import (
	"github.com/CreFire/rain/api"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/model"
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
	logger.Info("start rain")
	initDBTable()
	api.Router(r)

	// 运行服务器
	r.Run(":8080")
}

func initDBTable() {
	engine := dal.GetDb()
	err := engine.Sync2(new(model.User))
	if err != nil {
		log.Fatal("Could not synchronize database", zap.Error(err))
	}
	err = engine.Sync2(new(model.Tenant))
	if err != nil {
		log.Fatal("Could not synchronize database", zap.Error(err))
	}
	err = engine.Sync2(new(model.Permission))
	if err != nil {
		log.Fatal("Could not synchronize database", zap.Error(err))
	}
	user := &model.User{Name: "root"}
	ex, err := engine.Get(user)
	if err != nil {
		return
	}
	if !ex {
		user.PassWord = "admin"
		_, err := engine.Insert(user)
		if err != nil {
			return
		}
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
