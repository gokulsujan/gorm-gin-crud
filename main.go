package main

import (
	"crud-api-gorm/configuration"
	"crud-api-gorm/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configuration.PortConfiguration()
	configuration.DBConfiguration()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"message": "Hello"})
	})
	r.POST("/createUser", controller.UserCreation)
	r.GET("/users", controller.UserData)
	r.POST("/userUpdate/:id", controller.UserUpdate)
	r.POST("/userDelete/:id", controller.DeleteUser)

	r.Run()

}
