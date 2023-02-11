package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/13-bit/bulletinbird/birds"
	"github.com/13-bit/bulletinbird/botd"
	"github.com/13-bit/bulletinbird/config"
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
	emptyBotd := botd.Botd{
		Bird:        b,
		LastUpdated: ts,
	}

	f, err := os.Create(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	botdJson, _ := json.MarshalIndent(emptyBotd, "", "  ")

	f.Write(botdJson)
}

func main() {
	ResetBotd()
}
