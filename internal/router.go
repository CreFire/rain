package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//认证路由器
	authorized := r.Group("/")
	authorized.Use()
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
	}
	r.GET("/home/:id/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
}
