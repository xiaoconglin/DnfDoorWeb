package user_domain

import (
	"github.com/golang-jwt/jwt/v4"
	user_query "myweb/application/query"
	user_entities "myweb/domain/User/entities"
	user_token "myweb/infrastructure/encryption"
	"time"
)

type userServices struct {
}

var UserServices userServices

func (item *userServices) GetById(id uint) (userDo *user_entities.UserInfoDo) {
	do := user_entities.UserInfoDo{}
	do.GetById(id)
	return &do
}

func (item *userServices) CheckToken(tokenString string) (*user_entities.UserInfoDo, error) {
	claims, tokenErr := user_token.CheckToken(tokenString)
	if tokenErr != nil {
		return nil, tokenErr
	}
	do := user_entities.UserInfoDo{}
	do.GetById(claims.ID)
	return &do, nil
}

func (item *userServices) Login(query user_query.LoginInQuery) (interface{}, error) {
	userDo := user_entities.UserInfoDo{}
	userInfo, err := userDo.First(map[string]interface{}{"user_name": query.Username, "password": query.Password})
	if err != nil {
		return nil, err
	}
	claims := user_token.CustomClaims{
		ID:       userInfo.ID,
		NickName: userInfo.NickName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token, tokenErr := user_token.GenerateToken(claims)
	if err != nil {
		return nil, tokenErr
	}
	return token, nil

}
