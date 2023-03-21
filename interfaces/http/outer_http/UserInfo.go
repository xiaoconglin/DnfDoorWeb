package outer_http

import (
	"github.com/gin-gonic/gin"
	user_application "myweb/application"
	"net/http"
)

func ApiGetUserInfo(c *gin.Context) {
	user := user_application.GetUserInfo(1)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "请求成功",
		"code":   20000,
		"data": map[string]any{
			"name": user.Name,
		},
	})
}
