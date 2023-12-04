package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOverduePayments(c *gin.Context) {
	// 获取欠费信息逻辑
	c.JSON(http.StatusOK, gin.H{"message": "Overdue payments"})
}

func SendNotification(c *gin.Context) {
	// 发送提醒逻辑
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}
