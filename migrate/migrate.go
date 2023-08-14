package main

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"
)

func init() {
	configuration.PortConfiguration()
	configuration.DBConfiguration()
}

func main() {
	configuration.DB.AutoMigrate(&models.User{})
}
