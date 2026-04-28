package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guutar/todo-app-backend/internal/config"
	"github.com/guutar/todo-app-backend/internal/database"
)

func main() {
	// Load ENV Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Database Connection
	// db, err = database.Connect(cfg.DatabaseURL, cfg.DatabaseName)
	_, err = database.Connect(cfg.DatabaseURL, cfg.DatabaseName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Router Initializaion
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Todo App Backend is running well!",
		})
	})

	// Run server
	router.Run(":" + cfg.Port)
}
