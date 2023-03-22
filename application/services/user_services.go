package user_application

import (
	user_assemblers "myweb/application/assemblers"
	user_entities "myweb/domain/User/entities"
	user_domain "myweb/domain/User/services"
	dto_user "myweb/interfaces/dto"
)

func GetUserInfo(id uint) *user_entities.UserInfoDo {
	user := user_domain.UserServices.GetById(id)
	return user
}

func UserLogin(inDto *dto_user.LoginInDto) (*dto_user.LoginOutDto, error) {
	assemblers := user_assemblers.LoginInAssemblers{}
	query := assemblers.DtoToQuery(inDto)
	Token, err := user_domain.UserServices.Login(query)
	if err != nil {
		return nil, err
	}
	var outDto dto_user.LoginOutDto
	token, _ := Token.(string)
	outDto.Token = token
	return &outDto, nil
}
