package dao

import (
	"redrockCommerce/model"
)

func Check(user model.User) bool {
	db := Init()
	result := db.Table("redrock").Model(&model.User{}).Where("username = ? AND password = ?", user.Username, user.Password).Limit(1).Find(&user)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}
