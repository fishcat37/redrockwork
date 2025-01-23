package api

import (
	"redrockCommerce/dao"
	"redrockCommerce/model"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	var userId model.UserId
	if err := c.ShouldBindUri(&userId); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := dao.SelectId(userId.Id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status":  "10000",
		"info":    "success",
		"message": "查询成功",
		"data": gin.H{
			"user": gin.H{
				"username": user.Username,
				"password": user.Password,
			},
		},
	})

}
