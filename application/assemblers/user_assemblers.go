package user_assemblers

import (
	user_dto "myweb/apis/dto"
	user_query "myweb/application/query"
	user_entities "myweb/domain/User/entities"
)

type LoginInAssemblers struct {
}

func (item *LoginInAssemblers) DtoToQuery(inDto *user_dto.LoginInDto) (query user_query.LoginInQuery) {
	query.Username = inDto.Username
	query.Password = inDto.Password
	return
}

type UserInfoInAssemblers struct {
}

func (item *UserInfoInAssemblers) DoToDto(do *user_entities.UserInfoDo) *user_dto.UserInfoOutDto {
	var dto = user_dto.UserInfoOutDto{}
	dto.NickName = do.NickName
	return &dto
}
