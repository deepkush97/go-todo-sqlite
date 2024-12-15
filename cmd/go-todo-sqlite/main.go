package main

import (
	"log"
	"os"

	"github.com/deepkush97/go-todo-sqlite/internal/db"
	"github.com/deepkush97/go-todo-sqlite/internal/todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	err := godotenv.Load()
	if err == nil {
		log.Print(".env file loaded successfully")
	} else {
		log.Print("No .env file found, using system environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Environment variable PORT not set. Using default port: %s", port)
	} else {
		log.Printf("Using port from environment variable: %s", port)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	err = router.SetTrustedProxies([]string{})
	if err != nil {
		log.Printf("Failed to configure trusted proxies: %v", err)
	}

	todo.RegisterRoutes(router)

	if err := router.Run(":" + port); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
