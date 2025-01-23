package api

import (
	"redrockCommerce/dao"
	"redrockCommerce/model"
	"redrockCommerce/services"

	"github.com/gin-gonic/gin"
)

func Password(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claim, err := services.ParseToken(token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "token无效",
		})
		return
	}
	var password model.Password
	c.BindJSON(&password)
	user := model.User{claim.Username, claim.Password}
	a := dao.Check(user)
	if !a {
		c.JSON(400, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}
	if password.OldPassword == password.NewPassword {
		c.JSON(400, gin.H{
			"message": "新密码不能与旧密码相同",
		})
		return
	}
	err = dao.ChangePassword(user, password.NewPassword)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "修改密码失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "10000",
		"info":    "success",
		"message": "修改密码成功",
	})
}
