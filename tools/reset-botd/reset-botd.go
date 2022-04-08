package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/13-bit/birdboard/internal/botd"
	"github.com/13-bit/birdboard/internal/config"
)

func ResetBotd() {
	err := os.Remove(config.BotdFilePath())
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	birds.RegenerateBirdList()

	b := birds.Bird{}
	ts, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	emptyBotd := botd.BirdOfTheDay{
		Bird:        b,
		LastUpdated: ts,
	}

	botd.SaveBotd(emptyBotd)
}

func main() {
	ResetBotd()
}
