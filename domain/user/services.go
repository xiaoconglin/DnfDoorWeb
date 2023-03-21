package user_domain

type userServices struct {
}

var UserServices userServices

func (item *userServices) GetById(id uint) (userPo *UserInfoDo) {
	do := UserInfoDo{}
	userPo = do.GetById(id)
	return
}
