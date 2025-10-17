package main

import (
	"log"
	"os"
	"strconv"
	
	"discipline/api"
	"discipline/internal/handlers"
	"discipline/pkg/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	// Database configuration
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := database.InitDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize handlers
	userHandler := handlers.NewUserHandler(db)
	matchHandler := handlers.NewMatchHandler(db)
	roundHandler := handlers.NewRoundHandler(db)
	actionHandler := handlers.NewActionHandler(db)
	
	// Setup routes
	api.SetupRoutes(e, userHandler, matchHandler, roundHandler, actionHandler)
	
	// Simple test route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Discipline API is running!",
			"version": "1.0.0",
		})
	})

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	
	log.Printf("Server starting on port %s", serverPort)
	e.Logger.Fatal(e.Start(":" + serverPort))
}
