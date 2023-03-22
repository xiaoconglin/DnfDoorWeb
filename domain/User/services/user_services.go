package user_domain

import (
	user_query "myweb/application/query"
	user_entities "myweb/domain/User/entities"
)

type userServices struct {
}

var UserServices userServices

func (item *userServices) GetById(id uint) (userDo *user_entities.UserInfoDo) {
	do := user_entities.UserInfoDo{}
	userDo = do.GetById(id)
	return
}

func (item *userServices) Login(query user_query.LoginInQuery) (interface{}, error) {
	mapper := user_entities.UserInfoDoMapper{}
	_, err := mapper.First(map[string]interface{}{"user_name": query.Username, "password": query.Password})
	if err != nil {
		return nil, err
	}
	return "54321", nil

}
