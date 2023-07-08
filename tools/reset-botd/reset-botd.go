package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/13-bit/bulletinbird/birds"
	"github.com/13-bit/bulletinbird/botd"
	"github.com/13-bit/bulletinbird/config"
	"github.com/13-bit/bulletinbird/util"
)

func ResetBotd() {
	_ = os.Remove(config.BotdFilePath())

	file, err := os.Create(config.BotdFilePath())
	util.CheckError(err)

	defer file.Close()

	birds.RegenerateBirdList()

	b := birds.Bird{}
	ts, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	emptyBotd := botd.Botd{
		Bird:        b,
		LastUpdated: ts,
	}

	f, err := os.Create(config.BotdFilePath())
	util.CheckError(err)

	defer f.Close()

	botdJson, _ := json.MarshalIndent(emptyBotd, "", "  ")

	f.Write(botdJson)
}

func main() {
	ResetBotd()
}
