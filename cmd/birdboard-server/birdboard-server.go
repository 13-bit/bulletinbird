package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/13-bit/birdboard/internal/botd"
	"github.com/13-bit/birdboard/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/botd", birdOfTheDay)

	e.Static("/static", config.StaticPath())

	// Start server
	e.Logger.Fatal(e.Start(":1313"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Birdboard!")
}

func birdOfTheDay(c echo.Context) error {
	botdJson, err := json.Marshal(botd.GetBirdOfTheDay())
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, string(botdJson))
}
