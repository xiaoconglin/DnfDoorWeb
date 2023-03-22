package user_dto

type LoginInDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutDto struct {
	Token string `json:"token"`
}
