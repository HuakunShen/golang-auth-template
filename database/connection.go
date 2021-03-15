package database

import (
	"github.com/HuakunShen/golang-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Hacker971@/go_auth"), &gorm.Config{})

	if err != nil {
		panic(err)
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
