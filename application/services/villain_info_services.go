package application_services

import (
	"myweb/apis/dto"
	"myweb/application/assemblers"
	"myweb/domain/server_villain/services"
)

func GetVillainInfo(id uint) *dto.VillainInfoOutDto {
	do := villain_domain.VillainServices.GetById(id)
	assemblers := application_assemblers.VillainInfoInAssemblers{}
	outDto := assemblers.DoToDto(do)
	return outDto
}
