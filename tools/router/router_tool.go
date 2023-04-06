package router_tool

import "github.com/gin-gonic/gin"

type RouterStruct struct {
	// 方法
	Method string
	// 路径
	Path string
	// 对应的处理方法
	Handler gin.HandlerFunc
}

// InitRouter 初始化路由
func InitRouter(router *gin.RouterGroup, urlList []RouterStruct) {
	for i := 0; i < len(urlList); i++ {
		item := urlList[i]
		router.Handle(item.Method, item.Path, item.Handler)
	}
}
