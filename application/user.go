package user_application

import (
	user_domain "myweb/domain/user"
)

func GetUserInfo(id uint) *user_domain.UserInfoDo {
	user := user_domain.UserServices.GetById(id)
	return user
}
