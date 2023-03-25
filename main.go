package main

import (
	"github.com/CreFire/rain/config"
	"github.com/CreFire/rain/internal"
	"github.com/CreFire/rain/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	internal.Router(r)
	port := config.GlobalConfig.Server.Port
	logger := log.Default()
	logger.With(zap.String("port", port)).Info("ads")
	err := r.Run(port)
	if err != nil {
		log.Error("web", zap.String("port", port))
	}
}
