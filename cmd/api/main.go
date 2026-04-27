package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guutar/todo-app-backend/internal/config"
)

func main() {
	var err error

	// Load ENV Configuration
	var cfg *config.Config
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
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
