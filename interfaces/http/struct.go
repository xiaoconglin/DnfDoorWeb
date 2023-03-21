package http

import "github.com/gin-gonic/gin"

type RouterStruct struct {
	// 方法
	Method string
	// 路径
	Path string
	// 对应的处理方法
	Handler gin.HandlerFunc
}
