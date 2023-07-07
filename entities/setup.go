package entities

import "github.com/ortizdavid/golang-fiber-webapp/config"

func SetupMigrations() {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&User{})
}