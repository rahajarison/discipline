package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	
	"github.com/labstack/echo/v4"
	"discipline/pkg/database"
)

func main() {
	e := echo.New()
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	_, err := database.InitDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
