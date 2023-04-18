package outer_http

import (
	"github.com/gin-gonic/gin"
	"myweb/apis/dto"
	"myweb/apis/http_apis"
	application_services "myweb/application/services"
)

func ApiGetUserInfo(c *gin.Context) {
	var inDto dto.UserInfoInDto
	if err := c.ShouldBindQuery(&inDto); err != nil {
		return
	}
	user := application_services.GetUserInfo(inDto.UserId)
	http_apis.SuccessJsonRespond(c, user)
	return
}

func ApiCheckToken(c *gin.Context) {
	var inDto dto.ApiCheckInfoInDto
	if err := c.ShouldBindQuery(&inDto); err != nil {
		return
	}
	user, err := application_services.CheckToken(inDto.Token)
	if err != nil {
		http_apis.FailJsonRespond(c, err)
		return
	}
	http_apis.SuccessJsonRespond(c, user)
	return
}

func ApiUserLogin(c *gin.Context) {
	var inDto dto.LoginInDto
	if err := c.BindJSON(&inDto); err != nil {
		return
	}
	outDto, err := application_services.UserLogin(&inDto)
	if err != nil {
		http_apis.FailJsonRespond(c, err)
		return
	}
	http_apis.SuccessJsonRespond(c, outDto)
	return
}
