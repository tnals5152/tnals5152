package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("", func(ctx *gin.Context) {
		fmt.Println("Test")

		var api map[string]any

		err := ctx.Bind(&api)
		fmt.Println(err)
	})

	return r
}

// "POST / HTTP/1.1\r\nHost: 9881-121-163-158-16.jp.ngrok.io\r\nUser-Agent: GitHub-Hookshot/ede37db\r\nContent-Length: 7286\r\nAccept: */*\r\nContent-Type: application/json\r\nX-Forwarded-For: 140.82.115.96\r\nX-Forwarded-Proto: https\r\nX-Github-Delivery: 9afaa3d0-5813-11ed-8139-09a6076bec9d\r\nX-Github-Event: push\r\nX-Github-Hook-Id: 386164052\r\nX-Github-Hook-Installation-Target-Id: 471745865\r\nX-Github-Hook-Installation-Target-Type: repository\r\nAccept-Encoding: gzip\r\n\r\n{\"ref\":\"refs/heads/main\",\"before\":\"1f83d8bc99113d5640654df55"
