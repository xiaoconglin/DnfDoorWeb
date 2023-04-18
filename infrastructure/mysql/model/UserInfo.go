package mysql_model

import (
	"time"
)

type UserInfoRePo struct {
	ID             uint      `json:"id"`
	UserName       string    `json:"user_name"`
	Password       string    `json:"password"`
	NickName       string    `json:"nick_name"`
	CreateDatetime time.Time `json:"create_time"`
	UpdateDatetime time.Time `json:"update_time"`
}

func (UserInfoRePo) TableName() string {
	return "user_info"
}
