package outer_http

import (
	"github.com/gin-gonic/gin"
	user_application "myweb/application/services"
	user_dto "myweb/interfaces/dto"
	"net/http"
)

func ApiGetUserInfo(c *gin.Context) {
	user := user_application.GetUserInfo(1)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "请求成功",
		"code":   20000,
		"data": map[string]any{
			"name": user.NickName,
		},
	})
}

func ApiUserLogin(c *gin.Context) {
	var inDto user_dto.LoginInDto
	if err := c.BindJSON(&inDto); err != nil {
		return
	}
	outDto, err := user_application.UserLogin(&inDto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  errorToString(err),
			"data": nil,
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "请求成功",
		"code":   20000,
		"data": map[string]any{
			"token": outDto.Token,
		},
	})
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
