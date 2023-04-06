package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 定义登录、注册和认证路由器
	authRouter := r.Group("/auth")
	userRouter := r.Group("/user")

	// 注册路由处理函数
	r.POST("/register", registerEndpoint)

	// 登录路由处理函数
	authRouter.POST("/login", loginEndpoint)

	// 在认证路由器中使用中间件，验证是否已登录
	authRouter.Use(AuthMiddleware())
	{
		// 在用户路由器中定义需要进行权限验证的路由
		userRouter.GET("/profile", AllUserInfoEndpoint)
	}

}

// AuthMiddleware 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 请求头信息
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 如果 Authorization 为空，则返回未授权错误
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 验证 Authorization 是否合法，这里可以使用 JWT 鉴权等方式进行验证
		// ...

		// 如果验证通过，将用户信息保存到 Context 中
		c.Set("user_id", 123)

		c.Next()
	}
}
