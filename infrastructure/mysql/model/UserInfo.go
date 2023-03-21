package mysql_model

import (
	"time"
)

type UserInfoPo struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	CreateDatetime time.Time `json:"create_time"`
	UpdateDatetime time.Time `json:"update_time"`
}

func (UserInfoPo) TableName() string {
	return "user_info"
}
