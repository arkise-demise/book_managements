package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() (*gorm.DB, error) {
	d, err := gorm.Open("mysql", "root:root@/simpleRust?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db = d
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
