package internal

import (
	"github.com/CreFire/rain/internal/common"
	"github.com/CreFire/rain/internal/dal"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/tools/log"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	"xorm.io/xorm"
)

func loginEndpoint(c *gin.Context) {
	// Get the username and password from the request
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Get a new *xorm.Engine instance
	engine := dal.GetDb()
	defer func(engine *xorm.Engine) {
		err := engine.Close()
		if err != nil {
			log.Error("engine close err", zap.Error(err))
		}
	}(engine) // Close the engine at the end of the function

	// Create a new User object to search for
	user := &model.User{Email: username}

	// Use the xorm Engine's Get method to retrieve the User from the database
	if _, err := engine.Get(user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Check if the password is correct
	if user.PassWord != password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Password is correct, create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"role": user.Position,
		"time": time.Now().Unix(),
	})

	// Generate a signed token string
	tokenString, err := token.SignedString([]byte(common.MySecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	// Return the token string as JSON
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func SelectIdEndpoint(c *gin.Context) {
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
	user := &model.User{Id: uint(id)}

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

func registerEndpoint(c *gin.Context) {
	// 解析请求体中的 JSON 数据
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取数据库连接
	engine := dal.GetDb()
	defer engine.Close()

	// 将用户信息插入到数据库中
	if _, err := engine.Insert(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
