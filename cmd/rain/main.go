package main

import (
	"github.com/CreFire/rain/internal/server/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	routers.Router(engine)
}
