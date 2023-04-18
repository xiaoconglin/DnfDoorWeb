package dto

type UserInfoInDto struct {
	UserId uint `form:"user_id"`
}

type ApiCheckInfoInDto struct {
	Token string `form:"token"`
}

type LoginInDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
