package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dsn string

func Init() *gorm.DB { //初始化数据库连接
	// dbname := model.DbName{"root", "072231", "127.0.0.1", 3306, "first"}
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbname.Username, dbname.Password, dbname.Host, dbname.Port, dbname.Dbname)
	Dsn = "root:072231@tcp(127.0.0.1:3306)/first?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，error:" + err.Error())
	}
	return db
}
