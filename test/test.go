package api

import (
	"net/http"
	"time"

	"redrockCommerce/model"
	"redrockCommerce/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 定义 JWT 的密钥
var jwtSecret = []byte("your_secret_key")

// 解析和验证 token 的函数
func ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// 刷新 token 的函数
func RefreshToken(claims *model.CustomClaims) (string, error) {
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour)) // 1 小时后过期
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Refresh(c *gin.Context) {
	// 1. 获取请求头中的 token
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token为空"})
		return
	}

	// 2. 解析 token
	claims, err := ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token解析失败",
		})
		return
	}

	// 3. 刷新 token
	newAccessToken, err := RefreshToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "token刷新失败",
		})
		return
	}

	// 4. 生成新的刷新令牌
	newRefreshToken, err := services.CreateRefreshToken(model.User{
		Username: claims.Username,
		Password: claims.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "刷新令牌生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
