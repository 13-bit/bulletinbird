package botd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/13-bit/birdboard/internal/config"
)

type BirdOfTheDay struct {
	Bird        birds.Bird `json:"bird"`
	LastUpdated time.Time  `json:"lastUpdated"`
}

func nextBirdOfTheDay() BirdOfTheDay {
	birdList := birds.GetBirdList()

	front := birdList.Front()
	bird := front.Value.(birds.Bird)

	birdList.Remove(front)
	birds.SaveBirdList(birdList)

	botd := BirdOfTheDay{
		Bird:        bird,
		LastUpdated: time.Now(),
	}

	SaveBotd(botd)

	return botd
}

func SaveBotd(botd BirdOfTheDay) {
	f, err := os.Create(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	botdJson, err := json.MarshalIndent(botd, "", "  ")

	f.Write(botdJson)
}

func GetBirdOfTheDay() birds.Bird {
	f, err := os.Open(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var botd BirdOfTheDay

	err = json.NewDecoder(f).Decode(&botd)
	if err != nil {
		botd = nextBirdOfTheDay()
	}

	if !isTodaysBotd(botd.LastUpdated) {
		botd = nextBirdOfTheDay()
	}

	fmt.Printf("%+v\n", botd.Bird)

	return botd.Bird
}

func isTodaysBotd(botdTime time.Time) bool {
	loc, _ := time.LoadLocation("America/Chicago")
	year, month, day := time.Now().In(loc).Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)

	return botdTime.After(midnight)
}
