package middleware

import (
	"fmt"
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RoleMiddleware(requiredRole common.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "授权格式错误",
			})
			return
		}
		tokenString := parts[1]

		_, err := service.ParseToken(tokenString)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					// Token格式错误
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token过期或未激活
				} else {
					// Token签名无效
				}
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "无效Token"})
				return
			}
		}
		userID, _ := c.Get("userID")
		var user = model.User{}
		has, err := dal.GetDb().ID(userID).Get(&user)
		if err != nil || !has {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无法获取用户信息",
			})
			return
		}

		if *user.Role < int32(requiredRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "您没有权限访问此资源",
			})
			return
		}
		c.Next()
	}
}

// AuthMiddleware 验证中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 请求头信息
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "未提供Token",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return common.MySecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效Token",
			})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效Token",
			})
			return
		}

		// 将用户ID存入上下文中
		c.Set("userID", int64(claims["id"].(float64)))
		c.Next()
	}
}
