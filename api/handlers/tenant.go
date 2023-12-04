package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTenant(c *gin.Context) {
	// 获取租户信息逻辑
	c.JSON(http.StatusOK, gin.H{"message": "Tenant info"})
}

func PayRent(c *gin.Context) {
	// 租户支付逻辑
	c.JSON(http.StatusOK, gin.H{"message": "Rent paid"})
}
