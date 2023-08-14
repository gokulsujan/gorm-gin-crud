package configuration

import (
	"log"

	"github.com/joho/godotenv"
)

// Port Configuration from env file
func PortConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
}
