package http_apis

import (
	"github.com/gin-gonic/gin"
	"myweb/infrastructure/goerror"
	dto_interface "myweb/infrastructure/interface"
	"net/http"
)

type JsonRespondCode string
type JsonRespondMessage string

var (
	SuccessCode    JsonRespondCode    = "0000"
	SuccessMessage JsonRespondMessage = "OK"

	FailCode    JsonRespondCode    = "FFFF"
	FailMessage JsonRespondMessage = "系统异常"
)

func SuccessJsonRespond(c *gin.Context, data dto_interface.DtoInterface) {
	c.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"msg":  SuccessMessage,
		"data": data,
	})
	c.Abort()
}
func FailJsonRespond(c *gin.Context, err error) {
	gErr, ok := err.(*goerror.GoError)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": FailCode,
			"msg":  FailMessage,
			"tips": goerror.ErrorToString(err),
			"data": nil,
		})
		c.Abort()
		return
	}
	if gErr.Component == goerror.ErrApplication || gErr.Component == goerror.ErrInterfaces {
		c.JSON(http.StatusOK, gin.H{
			"code": gErr.Code,
			"msg":  gErr.Message,
			"data": nil,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":         FailCode,
		"msg":          FailMessage,
		"tips":         gErr.Message,
		"data":         nil,
		"causes":       gErr.Causes,
		"component":    gErr.Component,
		"responseType": gErr.ResponseType,
	})
	c.Abort()
	return
}
