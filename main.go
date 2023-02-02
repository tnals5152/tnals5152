package main

import (
	"net/http"
	"time"
	"fmt"
	"tnals5152/ssh/routers"
	"github.com/muesli/cache2go"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")

	routersInit := routers.InitRouter()
	readTimeout := 10 * time.Second
	writeTimeout := 10 * time.Second
	endPoint := ":" + "1995"
	maxHeaderBytes := 1 << 20
	fmt.Pritln(cache2go.Cache("cacheTable"))

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
}
