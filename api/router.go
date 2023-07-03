package api

import (
	"fmt"
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}
	r.POST("/register", registerHandler)
	r.GET("/select", SelectIdHandler)
	r.POST("/login", loginHandler)

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(authMiddleware(), roleMiddleware(common.ROLE_ADMIN))
	{
		adminRoutes.GET("/dashboard", adminDashboardHandler)
	}

	// 医生路由
	doctorRoutes := r.Group("/doctor")
	doctorRoutes.Use(authMiddleware(), roleMiddleware(common.ROLE_DICTOR))
	{
		doctorRoutes.GET("/patients", doctorPatientsHandler)
	}

	// 会员路由
	memberRoutes := r.Group("/member")
	memberRoutes.Use(authMiddleware(), roleMiddleware(common.ROLE_MEMBER))
	{
		memberRoutes.GET("/profile", memberProfileHandler)
	}
}

func roleMiddleware(requiredRole common.ROLE_TYPE) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		var user = model.User{}
		has, err := dal.GetDb().ID(userID).Get(&user)
		if err != nil || !has {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无法获取用户信息",
			})
			return
		}

		if *user.Role != int32(requiredRole) {
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
