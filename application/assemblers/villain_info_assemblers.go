package application_assemblers

import (
	"myweb/apis/dto"
	villain_entities "myweb/domain/server_villain/entities"
)

type VillainInfoInAssemblers struct {
}

func (item *VillainInfoInAssemblers) DoToDto(do *villain_entities.VillainInfoDo) *dto.VillainInfoOutDto {
	var outDto = dto.VillainInfoOutDto{}
	outDto.ID = do.ID
	outDto.Regional = do.Regional
	outDto.Server = do.Server
	outDto.Job = do.Job
	outDto.NickName = do.NickName
	outDto.Status = do.Status
	outDto.CreateDatetime = do.CreateDatetime
	outDto.UpdateDatetime = do.UpdateDatetime
	return &outDto
}
