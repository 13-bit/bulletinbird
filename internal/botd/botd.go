package botd

import (
	"encoding/json"
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
	return botd
}

func InitBotd() {
	botd = loadBotd()

	UpdateBotd()
}

func UpdateBotd() {
	lock.Lock()
	defer lock.Unlock()

	if !isTodaysBotd(botd.LastUpdated) {
		botd = nextBotd()
	}

	fmt.Printf("%+v\n", botd.Bird)

	refresh()
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

	saveBotd(botd)

	return botd
}

func refresh() {
	downloadBotdImage(botd)
	generateQrCode(botd, 80)
	generateQrCode(botd, 144)
	downloadLifeHistoryImages(botd)

	magtag.GenerateMagtagImages()
	inky.GenerateInkyImages(botd.Bird)
	inky.Refresh()
}

func saveBotd(botd Botd) {
	f, err := os.Create(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	botdJson, _ := json.MarshalIndent(botd, "", "  ")

	f.Write(botdJson)
}

func loadBotd() Botd {
	fmt.Println("Loading BOTD...")

	f, err := os.Open(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var b Botd

	err = json.NewDecoder(f).Decode(&b)
	if err != nil {
		log.Fatal(err)
	}

	return b
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
	qrPath := config.QrCodeImageDownloadPath(size)

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
