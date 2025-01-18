package main

import (
	"redrockCommerce/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.Router()
	r.Run(":8080")
}
