package api

import (
	"net/http"
	"redrockCommerce/model"
	"redrockCommerce/services"
	"time"

	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var refresh_token model.RefreshToken
	err := c.BindJSON(&refresh_token)
	if err != nil {

		c.JSON(400, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}
	claims1, err1 := services.ParseToken(token)
	claims2, err2 := services.ParseToken(refresh_token.Token)
	if err1 != nil || err2 != nil {
		c.JSON(401, gin.H{
			"message": "token无效",
		})
		return
	}
	if time.Now().Unix() > claims2.ExpiresAt.Time.Unix() {
		c.JSON(401, gin.H{
			"message": "refresh_token已过期",
		})
		return
	}
	info := model.User{
		Username: claims1.Username,
		Password: claims1.Password,
	}
	accessToken, err := services.CreateAccessToken(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成访问令牌失败"})
		return
	}
	refreshToken, err := services.CreateRefreshToken(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成刷新令牌失败"})
		return
	}
	c.JSON(200, gin.H{
		"status":        "10000",
		"info":          "success",
		"refresh_token": refreshToken,
		"token":         accessToken,
	})

}
