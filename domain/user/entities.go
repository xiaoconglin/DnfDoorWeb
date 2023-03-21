package user_domain

import (
	mysql_gorm "myweb/infrastructure/mysql"
	mysql_model "myweb/infrastructure/mysql/model"
	"time"
)

// do

type userInfoDoMapper struct{}

func (item *userInfoDoMapper) poToDo(po *mysql_model.UserInfoPo, do *UserInfoDo) {
	do.ID = po.ID
	do.Name = po.Name
	do.CreateDatetime = po.CreateDatetime
	do.UpdateDatetime = po.UpdateDatetime
	return
}
func (item *userInfoDoMapper) getById(id uint) (user mysql_model.UserInfoPo) {
	mysql_gorm.Db.First(&user, id)
	return
}

type UserInfoDo struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	CreateDatetime time.Time `json:"create_time"`
	UpdateDatetime time.Time `json:"update_time"`
	mapper         userInfoDoMapper
}

func (item UserInfoDo) GetById(id uint) *UserInfoDo {
	userPo := item.mapper.getById(id)
	item.mapper.poToDo(&userPo, &item)
	return &item
}
