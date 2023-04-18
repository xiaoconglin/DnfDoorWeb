package outer_http

import (
	"github.com/gin-gonic/gin"
	"myweb/apis/dto"
	"myweb/apis/http_apis"
	"myweb/application/services"
)

func ApiVillainInfo(c *gin.Context) {
	var inDto dto.VillainInfoInDto
	if err := c.ShouldBindQuery(&inDto); err != nil {
		return
	}
	outDto := application_services.GetVillainInfo(inDto.VillainInfoID)
	http_apis.SuccessJsonRespond(c, outDto)
	return
}
