package user_entities

import (
	mysql_gorm "myweb/infrastructure/mysql"
	mysql_model "myweb/infrastructure/mysql/model"
	"time"
)

type UserInfoDoMapper struct {
}

func (item *UserInfoDoMapper) PoToDo(po *mysql_model.UserInfoRePo, do *UserInfoDo) {
	do.ID = po.ID
	do.UserName = po.UserName
	do.Password = po.Password
	do.NickName = po.NickName
	do.CreateDatetime = po.CreateDatetime
	do.UpdateDatetime = po.UpdateDatetime
	return
}
func (item *UserInfoDoMapper) GetById(id uint) (user mysql_model.UserInfoRePo) {
	mysql_gorm.Db.First(&user, id)
	return
}

func (item *UserInfoDoMapper) Find(condition map[string]interface{}) *[]mysql_model.UserInfoRePo {
	var users []mysql_model.UserInfoRePo
	err := mysql_gorm.Db.Where(condition).Find(&users)
	if err != nil {
		return nil
	}
	return &users
}
func (item *UserInfoDoMapper) First(condition map[string]interface{}) (*mysql_model.UserInfoRePo, error) {
	var user mysql_model.UserInfoRePo
	err := mysql_gorm.Db.Where(condition).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type UserInfoDo struct {
	ID             uint      `json:"id"`
	UserName       string    `json:"user_name"`
	Password       string    `json:"password"`
	NickName       string    `json:"nick_name"`
	CreateDatetime time.Time `json:"create_time"`
	UpdateDatetime time.Time `json:"update_time"`
	mapper         UserInfoDoMapper
}

func (item *UserInfoDo) GetById(id uint) {
	userPo := item.mapper.GetById(id)
	item.mapper.PoToDo(&userPo, item)
	return
}

func (item *UserInfoDo) First(condition map[string]interface{}) (*UserInfoDo, error) {
	userPo, err := item.mapper.First(condition)
	if err != nil {
		return nil, err
	}
	item.mapper.PoToDo(userPo, item)
	return item, nil
}
