package dto

import "time"

type VillainInfoOutDto struct {
	OutDto
	ID             uint      `json:"id"`
	Regional       string    `json:"regional"`
	Server         string    `json:"server"`
	Job            string    `json:"job"`
	Role           string    `json:"role"`
	NickName       string    `json:"nick_name"`
	Status         string    `json:"status"`
	CreateDatetime time.Time `json:"create_time"`
	UpdateDatetime time.Time `json:"update_time"`
}
