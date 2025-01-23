package dao

import "redrockCommerce/model"

func SelectId(id int) (model.User, error) {
	db := Init()
	var user model.User
	err := db.Table("redrock").Select("username", "password").Where("id=?", id).Find(&user).Error
	return user, err
}
