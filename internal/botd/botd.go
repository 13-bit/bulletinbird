package botd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/13-bit/bulletinbird-server/internal/birds"
	"github.com/13-bit/bulletinbird-server/internal/config"
	"github.com/13-bit/bulletinbird-server/internal/platforms/inky"
	"github.com/13-bit/bulletinbird-server/internal/platforms/magtag"
	"github.com/cavaliergopher/grab/v3"
	qrcode "github.com/skip2/go-qrcode"
)

var lock = &sync.Mutex{}

type Botd struct {
	Bird        birds.Bird `json:"bird"`
	LastUpdated time.Time  `json:"lastUpdated"`
}

var (
	botd Botd
)

func BirdOfTheDay() Botd {

	lock.Lock()
	defer lock.Unlock()

	if !isTodaysBotd(botd.LastUpdated) {
		LoadBotd()
	}

	return botd
}

func UpdateBotd() error {
	lock.Lock()
	defer lock.Unlock()

	if !isTodaysBotd(botd.LastUpdated) {
		botd = nextBotd()
	} else {
		return errors.New("BOTD is already up to date")
	}

	fmt.Printf("%+v\n", botd.Bird)

	return nil
}

func nextBotd() Botd {
	birdList := birds.GetBirdList()

	front := birdList.Front()
	bird := front.Value.(birds.Bird)

	birdList.Remove(front)
	birds.SaveBirdList(birdList)

	botd := Botd{
		Bird:        bird,
		LastUpdated: time.Now(),
	}

	SaveBotd(botd)
	downloadBotdImage(botd)
	generateQrCode(botd, 80)
	downloadLifeHistoryImages(botd)

	magtag.GenerateMagtagImages()
	inky.GenerateInkyImages()

	return botd
}

func SaveBotd(botd Botd) {
	f, err := os.Create(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	botdJson, _ := json.MarshalIndent(botd, "", "  ")

	f.Write(botdJson)
}

func LoadBotd() {
	f, err := os.Open(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = json.NewDecoder(f).Decode(&botd)
	if err != nil {
		botd = nextBotd()
	}

	if !isTodaysBotd(botd.LastUpdated) {
		botd = nextBotd()
	}

	fmt.Printf("%+v\n", botd.Bird)
}

func isTodaysBotd(botdTime time.Time) bool {
	loc, _ := time.LoadLocation("America/Chicago")
	year, month, day := time.Now().In(loc).Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)

	return botdTime.After(midnight)
}

func downloadBotdImage(botd Botd) {
	fmt.Printf("Downloading BOTD image from %s...\n", botd.Bird.IllustrationUrl)

	botdPath := config.BotdImageDownloadPath()

	err := os.Remove(botdPath)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := grab.Get(botdPath, botd.Bird.IllustrationUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)
}

func generateQrCode(botd Botd, size int) {
	qrPath := config.QrCodeImageDownloadPath()

	err := qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Medium, size, qrPath)
	if err != nil {
		fmt.Println(err)
	}
}

func downloadLifeHistoryImages(botd Botd) {
	fmt.Println("Downloading life history images...")

	habitatPath, foodPath, nestingPath, behaviorPath, conservationPath := config.LifeHistoryImageDownloadPaths()

	lifeHistoryPaths := []string{habitatPath, foodPath, nestingPath, behaviorPath, conservationPath}

	for index, path := range lifeHistoryPaths {
		err := os.Remove(path)
		if err != nil {
			fmt.Println(err)
		}

		resp, err := grab.Get(path, botd.Bird.LifeHistoryImageUrls[index])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Download saved to", resp.Filename)
	}
}
