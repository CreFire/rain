package service

import (
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/utils/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func LoginHandler(c *gin.Context) {
	// Get the username and password from the request
	user := model.User{}
	// Create a new User object to search for
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("login 1", zap.Any("user", user))
	// Get a new *xorm.Engine instance
	engine := dal.GetDb()
	curUser := model.User{
		Name: user.Name,
	}
	// Use the xorm Engine's Get method to retrieve the User from the database
	if _, err := engine.Get(&curUser); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Check if the password is correct
	if curUser.PassWord != user.PassWord {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Password is correct, create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
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
		"message": "登录成功",
		"token":   tokenString,
	})
}
