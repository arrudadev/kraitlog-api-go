package main

import (
	"log"

	"github.com/arrudadev/kraitlog-api/config"
	userRoutes "github.com/arrudadev/kraitlog-api/internal/infrastructure/api/user/routes"
	"github.com/arrudadev/kraitlog-api/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.NewConnection(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	userRoutes.RegisterUserRoutes(router, db)

	router.Run(":8080")
}
