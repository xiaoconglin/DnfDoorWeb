package villain_domain

import (
	villain_entities "myweb/domain/server_villain/entities"
)

type villainServices struct {
}

var VillainServices villainServices

func (item *villainServices) GetById(id uint) *villain_entities.VillainInfoDo {
	do := villain_entities.VillainInfoDo{}
	do.GetById(id)
	return &do
}
