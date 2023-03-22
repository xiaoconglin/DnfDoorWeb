package main

import (
	"github.com/gin-gonic/gin"
	"myweb/interfaces/http/outer_http"
	"myweb/middlewares"
	router_tool "myweb/tools"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.Core())

	outer_router := router.Group("outer_api")
	router_tool.InitRouter(outer_router, outer_http.OuterUrlList)
	router.Run() // listen and serve on 0.0.0.0:8080
}
