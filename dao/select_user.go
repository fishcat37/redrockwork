package dao

import (
	"redrockCommerce/model"
)

func SelecrUser(name string) bool { //查询用户是否存在
	DB := Init()
	var user model.User
	DB.Table("redrock").Where("username=?", name).Find(&user)
	if user.Username == name {
		return true
	}
	return false
}
