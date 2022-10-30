package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("webhook", func(ctx *gin.Context) {
		fmt.Println("Test")
	})

	return r
}
