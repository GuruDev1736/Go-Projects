package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB // create the db instance
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Guruprasad1736@tcp(localhost:3306)/stock?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB { // return the db instance
	return db
}
