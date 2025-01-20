package dao

import (
	"redrockCommerce/model"
)

func AddUser(req model.User) error { //注册,将用户信息存入数据库
	DB := Init()
	err := DB.Table("redrock").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}
