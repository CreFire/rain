package handlers

import (
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/service"
	"github.com/CreFire/rain/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LoginHandler(c *gin.Context) {
	var loginReq model.LoginRequest
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 假设 GetUserByUsername 是从数据库获取用户信息的函数
	user, err := model.GetUserByAccount(loginReq.Account)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// 检查密码
	if !user.CheckPassword(loginReq.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// 生成JWT令牌
	token, err := service.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token string as JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

func SelectIdHandler(c *gin.Context) {
	// Get the ID parameter from the request URL
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	// Create a new xorm Engine and connect to the database
	engine := dal.GetDb()
	engine.Close()

	// Create a new User object to search for
	user := &model.User{Id: &id}

	// Use the xorm Engine's Get method to retrieve the User from the database
	if _, err := engine.Get(user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Return the found User as JSON
	c.JSON(http.StatusOK, user)
}

func RegisterHandler(c *gin.Context) {
	// 解析请求体中的 JSON 数据
	var user = &model.User{}
	if err := c.BindJSON(user); err != nil {
		log.Error("bad req", "err", err, "user", user)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 数据处理
	code, obj := service.UpdateUser(user)
	// 返回响应
	c.JSON(code, obj)
}
