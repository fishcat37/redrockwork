package api

import "github.com/gin-gonic/gin"

func Refresh(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(400, gin.H{"error": "未携带token"})
		return
	}
}
