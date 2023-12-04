package common

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func JSONResponse(c *gin.Context, statusCode int, success bool, data any, errMsg string) {
	resp := Response{
		Success: success,
		Data:    data,
		Error:   errMsg,
	}

	c.JSON(statusCode, resp)
}
