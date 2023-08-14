package controller

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func UserData(c *gin.Context) {
	var userList []models.User
	configuration.DB.Find(&userList)
	c.JSON(200, gin.H{"Users": userList})
}
