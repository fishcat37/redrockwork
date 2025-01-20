package api

import (
	"github.com/gin-gonic/gin"
)

func Router() { //Router是用来创建路由的函数，路由是用来处理客户端请求的函数
	r := gin.Default()
	r.POST("/user/register", Register)
	r.GET("/user/login", Token)
	r.GET("/user/token/refresh", Refresh)
	r.Run(":8080")
}
