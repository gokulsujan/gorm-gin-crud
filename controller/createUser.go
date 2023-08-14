package controller

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func UserCreation(c *gin.Context) {
	// assigning user details
	user := models.User{
		Name: "Gokul Sujan",
		Age:  24,
	}

	//inserting into user table
	result := configuration.DB.Create(&user)

	//If error occured
	if result.Error != nil {
		c.Status(200)
		return
	}

	c.JSON(200, gin.H{"Insereted": user})

}
