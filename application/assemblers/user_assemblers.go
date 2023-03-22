package user_assemblers

import (
	user_query "myweb/application/query"
	dto_user "myweb/interfaces/dto"
)

type LoginInAssemblers struct {
}

func (item *LoginInAssemblers) DtoToQuery(inDto *dto_user.LoginInDto) (query user_query.LoginInQuery) {
	query.Username = inDto.Username
	query.Password = inDto.Password
	return
}
