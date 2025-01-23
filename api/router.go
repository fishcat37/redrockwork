package api

import (
	"github.com/gin-gonic/gin"
)

func Router() { //Router是用来创建路由的函数，路由是用来处理客户端请求的函数
	r := gin.Default()
	r.POST("/user/register", Register)
	r.POST("/user/token", Token)
	r.GET("/user/token/refresh", Refresh)
	r.POST("/user/password", Password)
	r.GET("/user/info/:user_id", Info)
	r.Run(":8080")
}
