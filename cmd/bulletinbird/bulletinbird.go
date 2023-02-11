package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/13-bit/bulletinbird/botd"
	"github.com/13-bit/bulletinbird/config"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load the BOTD
	botd.InitBotd()

	// Update the BOTD every day at midnight
	loc, _ := time.LoadLocation("America/Chicago")
	s := gocron.NewScheduler(loc)

	s.Every(1).Day().At("00:01").Do(func() {
		botd.UpdateBotd()
	})

	s.StartAsync()

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
	return c.String(http.StatusOK, "Welcome to BulletinBird!")
}

func birdOfTheDay(c echo.Context) error {
	botdJson, err := json.Marshal(botd.BirdOfTheDay())
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, string(botdJson))
}
