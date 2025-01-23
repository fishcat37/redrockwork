package dao

import (
	"redrockCommerce/model"
)

func ChangePassword(user model.User, password string) error {
	db := Init()
	err := db.Table("redrock").Where("username=?", user.Username).Find(&user).Error
	if err != nil {
		return err
	}
	user.Password = password
	err = db.Table("redrock").Where("username = ?", user.Username).Update("password", password).Error
	if err != nil {
		return err
	}
	return nil
}
