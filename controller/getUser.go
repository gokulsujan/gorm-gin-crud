package controller

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func UserData(c *gin.Context) {
	result := configuration.DB.Find(&models.User{})

	if result.Error != nil {
		c.Status(500)
		return
	}
	c.JSON(200, gin.H{"Users": result})
}
