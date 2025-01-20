package api

import (
	"net/http"
	"redrockCommerce/dao"
	"redrockCommerce/model"
	"redrockCommerce/services"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	var info model.User
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	flag := dao.Check(info)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
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
