package outer_http

import (
	"github.com/gin-gonic/gin"
	user_dto "myweb/apis/dto"
	"myweb/apis/http_apis"
	user_application "myweb/application/services"
)

func ApiGetUserInfo(c *gin.Context) {
	var inDto user_dto.UserInfoInDto
	if err := c.ShouldBindQuery(&inDto); err != nil {
		return
	}
	user := user_application.GetUserInfo(inDto.UserId)
	http_apis.SuccessJsonRespond(c, user)
	return
}

func ApiCheckToken(c *gin.Context) {
	var inDto user_dto.ApiCheckInfoInDto
	if err := c.ShouldBindQuery(&inDto); err != nil {
		return
	}
	user, err := user_application.CheckToken(inDto.Token)
	if err != nil {
		http_apis.FailJsonRespond(c, err)
		return
	}
	http_apis.SuccessJsonRespond(c, user)
	return
}

func ApiUserLogin(c *gin.Context) {
	var inDto user_dto.LoginInDto
	if err := c.BindJSON(&inDto); err != nil {
		return
	}
	outDto, err := user_application.UserLogin(&inDto)
	if err != nil {
		http_apis.FailJsonRespond(c, err)
		return
	}
	http_apis.SuccessJsonRespond(c, outDto)
	return
}
