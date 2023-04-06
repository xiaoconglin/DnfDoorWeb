package user_dto

type OutDto struct {
}

func (item OutDto) ToDict() {
	return
}

type UserInfoInDto struct {
	UserId uint `form:"user_id"`
}

type UserInfoOutDto struct {
	OutDto
	NickName string `json:"nickName"`
}

type ApiCheckInfoInDto struct {
	Token string `form:"token"`
}

type LoginInDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutDto struct {
	OutDto
	Token string `json:"token"`
}
