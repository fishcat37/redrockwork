package api

import (
	"net/http"
	"redrockCommerce/dao"
	"redrockCommerce/model"
	"regexp"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "10000",
		"info":   "success",
	})
	name := req.Username
	password := req.Password
	a, err := regexp.MatchString("[a-zA-Z1-9]", password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if len(name) > 20 {
		c.JSON(http.StatusOK, gin.H{
			"error": "name must less than 20",
		})
		return
	} else if a == false || len(password) < 6 || len(password) > 20 {
		c.JSON(http.StatusOK, gin.H{
			"error": "password must contain number and letter and it must be 6-20",
		})
		return
	}
	err = dao.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
