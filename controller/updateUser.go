package controller

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		Age  int
	}

	c.BindJSON(&body)
	var post models.User
	configuration.DB.First(&post, id)
	configuration.DB.Model(&post).Updates(models.User{Name: body.Name, Age: body.Age})
	c.JSON(200, gin.H{"message": "success"})
}
