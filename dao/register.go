package dao

import (
	"redrockCommerce/model"
)

func Register(req model.RegisterRequest) error {
	DB := Init()
	err := DB.Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}
