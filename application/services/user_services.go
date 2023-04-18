package application_services

import (
	user_dto "myweb/apis/dto"
	user_assemblers "myweb/application/assemblers"
	user_domain "myweb/domain/User/services"
)

func GetUserInfo(id uint) *user_dto.UserInfoOutDto {
	user := user_domain.UserServices.GetById(id)
	assemblers := user_assemblers.UserInfoInAssemblers{}
	userDto := assemblers.DoToDto(user)
	return userDto
}

func UserLogin(inDto *user_dto.LoginInDto) (*user_dto.LoginOutDto, error) {
	assemblers := user_assemblers.LoginInAssemblers{}
	query := assemblers.DtoToQuery(inDto)
	Token, err := user_domain.UserServices.Login(query)
	if err != nil {
		return nil, err
	}
	var outDto user_dto.LoginOutDto
	token, _ := Token.(string)
	outDto.Token = token
	return &outDto, nil
}
func CheckToken(tokenString string) (*user_dto.UserInfoOutDto, error) {
	user, err := user_domain.UserServices.CheckToken(tokenString)
	if err != nil {
		return nil, err
	}
	assemblers := user_assemblers.UserInfoInAssemblers{}
	userDto := assemblers.DoToDto(user)
	return userDto, err
}
