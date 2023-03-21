package main

import (
	"github.com/gin-gonic/gin"
	"myweb/interfaces/http"
	"myweb/interfaces/http/outer_http"
	"myweb/middlewares"
)

// InitRouter 初始化路由
func InitRouter(router *gin.RouterGroup, urlList []http.RouterStruct) {
	for i := 0; i < len(urlList); i++ {
		item := urlList[i]
		router.Handle(item.Method, item.Path, item.Handler)
	}
}

func main() {
	router := gin.Default()
	router.Use(middlewares.Core())

	outer_router := router.Group("outer_api")
	InitRouter(outer_router, outer_http.OuterUrlList)
	router.Run() // listen and serve on 0.0.0.0:8080
}
