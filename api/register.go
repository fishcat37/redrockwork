package api

import (
	"net/http"
	"redrockCommerce/dao"
	"redrockCommerce/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) { //注册处理函数
	var info model.User
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := info.Username
	flag := dao.SelecrUser(name)
	if flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}
	id, err := dao.AddUser(info)
	if err != nil || id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "10000",
		"info":    "success",
		"message": "注册成功",
		"id":      id,
	})
}
