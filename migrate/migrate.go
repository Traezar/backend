package main

import (
	"fmt"
	"log"

	"github.com/Traezar/go-backend/initializers"
	"github.com/Traezar/go-backend/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec(" CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	initializers.DB.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{})
	fmt.Println("? Migration complete")
}
