package service

import (
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/dal"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/utils/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func UpdateUser(user *model.User) (code int, res *gin.H) {
	//验证数据

	// 获取数据库连接
	engine := dal.GetDb()
	// 将用户信息插入到数据库中
	if _, err := engine.Insert(user); err != nil {
		log.Error("bad req", "err", err, "user", user)
		code = http.StatusInternalServerError
		res = &gin.H{"error": err.Error()}
		return
	}
	code = http.StatusOK
	res = &gin.H{}
	return
}

// ParseToken 解析和验证 JWT
func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保JWT的签名方法是您期望的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(common.MySecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}

	// 验证令牌是否已过期
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
		}
	}

	return &claims, nil
}

// GenerateToken 生成JWT令牌，根据实际情况进行实现
func GenerateToken(user *model.User) (string, error) {
	expTime := time.Now().Add(time.Hour * 72).Unix() // 设置令牌有效期为72小时
	// Password is correct, create a new JWT token
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
		"time": time.Now().Unix(),
		"exp":  expTime,
	}).SignedString([]byte(common.MySecretKey))

	// Generate a signed token string
	if err != nil {
		slog.Error("err", "error", "Failed to generate token")
		return "", err
	}
	return tokenString, nil
}
