package controller

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	configuration.DB.Delete(&models.User{}, id)
	c.JSON(200, gin.H{"message": "User deleted"})
}
