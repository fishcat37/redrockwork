package dao

import (
	"redrockCommerce/model"
)

func AddUser(req model.User) (int, error) { //注册,将用户信息存入数据库
	db := Init()
	var id int
	err := db.Table("redrock").Create(&req).Pluck("id", &id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}
