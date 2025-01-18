package dao

import (
	"fmt"
	"redrockCommerce/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbname := model.DbName{"root", "072231", "127.0.0.1", 3306, "first"}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbname.Username, dbname.Password, dbname.Host, dbname.Port, dbname.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，error:" + err.Error())
	}
	return db
}
