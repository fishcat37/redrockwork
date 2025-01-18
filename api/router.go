package api

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine { //Router是用来创建路由的函数，路由是用来处理客户端请求的函数
	r := gin.Default()
	r.POST("/user/register", RegisterHandler)
	return r
}
