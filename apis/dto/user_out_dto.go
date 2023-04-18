package dto

type UserInfoOutDto struct {
	OutDto
	NickName string `json:"nickName"`
}

type LoginOutDto struct {
	OutDto
	Token string `json:"token"`
}
