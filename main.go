package main

import (
	"crud-api-gorm/configuration"
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

	r.Run()

}
