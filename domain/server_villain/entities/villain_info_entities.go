package villain_entities

import (
	mysql_gorm "myweb/infrastructure/mysql"
	mysql_model "myweb/infrastructure/mysql/model"
	"time"
)

type VillainInfoDoMapper struct {
}

func (item *VillainInfoDoMapper) PoToDo(repo *mysql_model.ServerVillainRePo, do *VillainInfoDo) {
	do.ID = repo.ID
	do.Regional = repo.Regional
	do.Server = repo.Server
	do.Job = repo.Job
	do.NickName = repo.NickName
	do.Status = repo.Status
	do.CreateDatetime = repo.CreateDatetime
	do.UpdateDatetime = repo.UpdateDatetime
	return
}
func (item *VillainInfoDoMapper) GetById(id uint) (user mysql_model.ServerVillainRePo) {
	mysql_gorm.Db.First(&user, id)
	return
}

func (item *VillainInfoDoMapper) Find(condition map[string]interface{}) *[]mysql_model.ServerVillainRePo {
	var repos []mysql_model.ServerVillainRePo
	err := mysql_gorm.Db.Where(condition).Find(&repos)
	if err != nil {
		return nil
	}
	return &repos
}
func (item *VillainInfoDoMapper) First(condition map[string]interface{}) (*mysql_model.ServerVillainRePo, error) {
	var repo mysql_model.ServerVillainRePo
	err := mysql_gorm.Db.Where(condition).First(&repo).Error
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

type VillainInfoDo struct {
	ID             uint      `json:"id"`
	Regional       string    `json:"regional"`
	Server         string    `json:"server"`
	Job            string    `json:"job"`
	Role           string    `json:"role"`
	NickName       string    `json:"nick_name"`
	Status         string    `json:"status"`
	CreateDatetime time.Time `json:"create_time"`
	UpdateDatetime time.Time `json:"update_time"`
	mapper         VillainInfoDoMapper
}

func (item *VillainInfoDo) GetById(id uint) *VillainInfoDo {
	rePo := item.mapper.GetById(id)
	item.mapper.PoToDo(&rePo, item)
	return item
}

func (item *VillainInfoDo) First(condition map[string]interface{}) (*VillainInfoDo, error) {
	repo, err := item.mapper.First(condition)
	if err != nil {
		return nil, err
	}
	item.mapper.PoToDo(repo, item)
	return item, nil
}
