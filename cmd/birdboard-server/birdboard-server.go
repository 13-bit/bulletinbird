package main

import (
	"fmt"

	"github.com/13-bit/birdboard/internal/botd"
)

type Configuration struct {
	EbirdToken string
}

func main() {
	fmt.Println("Welcome to Birdboard!")

	botd.BirdOfTheDay()
}
