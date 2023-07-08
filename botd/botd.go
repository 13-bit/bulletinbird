package botd

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/13-bit/bulletinbird/birds"
	"github.com/13-bit/bulletinbird/config"
	"github.com/13-bit/bulletinbird/platforms/inky"
	"github.com/13-bit/bulletinbird/platforms/magtag"
	"github.com/13-bit/bulletinbird/util"
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

func init() {
	birdlistFileExists, botdFileExists := checkDataFilesExist()

	if !birdlistFileExists {
		log.Printf("%s does not exist, creating...\n", config.BirdListFilePath())
		createBirdlistFile()
	}

	if !botdFileExists {
		log.Printf("%s does not exist, creating...\n", config.BotdFilePath())
		createBotdFile()
	}
}

func checkDataFilesExist() (bool, bool) {
	birdlistFileExists, botdFileExists := true, true

	_, err := os.Stat(config.BirdListFilePath())
	if os.IsNotExist(err) {
		birdlistFileExists = false
	}

	_, err = os.Stat(config.BotdFilePath())
	if os.IsNotExist(err) {
		botdFileExists = false
	}

	return birdlistFileExists, botdFileExists
}

func createBotdFile() {
	nextBotd()
}

func createBirdlistFile() {
	birds.RegenerateBirdList()
}

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

	log.Printf("%+v\n", botd.Bird)

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
}

func saveBotd(botd Botd) {
	f, err := os.Create(config.BotdFilePath())
	util.CheckError(err)

	defer f.Close()

	botdJson, _ := json.MarshalIndent(botd, "", "  ")

	f.Write(botdJson)
}

func loadBotd() Botd {
	log.Println("Loading BOTD...")

	f, err := os.Open(config.BotdFilePath())
	util.CheckError(err)

	defer f.Close()

	var b Botd

	err = json.NewDecoder(f).Decode(&b)
	util.CheckError(err)

	return b
}

func isTodaysBotd(botdTime time.Time) bool {
	loc, _ := time.LoadLocation("America/Chicago")
	year, month, day := time.Now().In(loc).Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, loc)

	return botdTime.After(midnight)
}

func downloadBotdImage(botd Botd) {
	log.Printf("Downloading BOTD image from %s...\n", botd.Bird.IllustrationUrl)

	botdPath := config.BotdImageDownloadPath()

	_ = os.Remove(botdPath)

	resp, err := grab.Get(botdPath, botd.Bird.IllustrationUrl)
	util.CheckError(err)

	log.Println("Download saved to", resp.Filename)
}

func generateQrCode(botd Botd, size int) {
	qrPath := config.QrCodeImageDownloadPath(size)

	err := qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Medium, size, qrPath)
	util.CheckError(err)
}

func downloadLifeHistoryImages(botd Botd) {
	log.Println("Downloading life history images...")

	habitatPath, foodPath, nestingPath, behaviorPath, conservationPath := config.LifeHistoryImageDownloadPaths()

	lifeHistoryPaths := []string{habitatPath, foodPath, nestingPath, behaviorPath, conservationPath}

	for index, path := range lifeHistoryPaths {
		_ = os.Remove(path)

		resp, err := grab.Get(path, botd.Bird.LifeHistoryImageUrls[index])
		util.CheckError(err)

		log.Println("Download saved to", resp.Filename)
	}
}
