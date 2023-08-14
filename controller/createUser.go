package controller

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func UserCreation(c *gin.Context) {
	// assigning user details

	user := models.User{}
	c.ShouldBindJSON(&user)

	//inserting into user table
	result := configuration.DB.Create(&user)

	//If error occured
	if result.Error != nil {
		c.Status(200)
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully inserted the data into database"})

}
