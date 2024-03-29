package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-auth/react-golang/models"
)

var DB *gorm.DB

func Connect() {
	connection,err := gorm.Open(mysql.Open("root:root@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}