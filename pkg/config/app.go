package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database_conn := "akshatphumbhra:password@tcp(devicetracker-mysql.cubueqoyn6uj.us-west-1.rds.amazonaws.com:3306)/devicetracker?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", database_conn)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
