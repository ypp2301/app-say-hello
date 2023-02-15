package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:80)/go_restapi_gin"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
	DB = database
}
