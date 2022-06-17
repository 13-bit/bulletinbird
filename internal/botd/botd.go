package botd

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"log"
	"os"
	"sync"
	"time"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/13-bit/birdboard/internal/config"
	"github.com/13-bit/birdboard/internal/img"
	"github.com/cavaliergopher/grab/v3"
	"github.com/disintegration/imaging"
	qrcode "github.com/skip2/go-qrcode"
)

//go:embed resources/*
var resourcesFS embed.FS

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
	processBotdImage()
	generateQrCode(botd)
	processQrCodeImage()
	downloadLifeHistoryImages(botd)
	processLifeHistoryImages()

	return botd
}

func SaveBotd(botd Botd) {
	f, err := os.Create(config.BotdFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	botdJson, err := json.MarshalIndent(botd, "", "  ")

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

func processBotdImage() {
	fmt.Println("Processing image...")

	botdPath := config.BotdImageDownloadPath()
	pngPath, bmpPath := config.BotdImageFilePaths()

	botdImage, err := imaging.Open(botdPath)
	if err != nil {
		log.Fatal(err)
	}

	// maskFile, err := resourcesFS.Open("resources/botd-mask.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer maskFile.Close()

	// maskImage, _, err := image.Decode(maskFile)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if botdImage.Bounds().Dx() > botdImage.Bounds().Dy() {
		botdImage = imaging.Resize(botdImage, 100, 0, imaging.Box)
	} else {
		botdImage = imaging.Resize(botdImage, 0, 80, imaging.Box)
	}

	// palette := []color.Color{
	// 	color.RGBA{76, 76, 76, 255},
	// 	color.RGBA{132, 132, 132, 255},
	// 	color.RGBA{221, 221, 221, 255},
	// 	color.RGBA{188, 188, 188, 255},
	// }

	// d := dither.NewDitherer(palette)
	// d.Matrix = dither.Atkinson

	// botdImageDithered := d.Dither(botdImage)

	// if botdImageDithered != nil {
	// 	botdImage = botdImageDithered
	// }

	// botdImage = imaging.Overlay(botdImage, maskImage, image.Pt(0, 0), 255)

	botdImage = img.RgbaToGray(botdImage)

	err = imaging.Save(botdImage, pngPath)
	if err != nil {
		fmt.Println(err)
	}

	err = imaging.Save(botdImage, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}

func generateQrCode(botd Botd) {
	qrPath := config.QrCodeImageDownloadPath()

	err := qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Medium, 80, qrPath)
	if err != nil {
		fmt.Println(err)
	}
}

func processQrCodeImage() {
	fmt.Println("Processing QR code...")

	qrPath := config.QrCodeImageDownloadPath()

	pngPath, bmpPath := config.QrCodeImageFilePaths()

	qrImage, err := imaging.Open(qrPath)
	if err != nil {
		log.Fatal(err)
	}

	qrImage = img.RgbaToGray(qrImage)

	err = imaging.Save(qrImage, pngPath)
	if err != nil {
		fmt.Println(err)
	}

	err = imaging.Save(qrImage, bmpPath)
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

func processLifeHistoryImages() {
	fmt.Println("Processing life history images...")

	lifeHistoryFile, err := resourcesFS.Open("resources/life-history-template.png")
	if err != nil {
		log.Fatal(err)
	}
	defer lifeHistoryFile.Close()

	lifeHistoryImage, _, err := image.Decode(lifeHistoryFile)
	if err != nil {
		fmt.Println(err)
	}

	habitatPath, foodPath, nestingPath, behaviorPath, _ := config.LifeHistoryImageDownloadPaths()

	habitatImage, err := imaging.Open(habitatPath)
	if err != nil {
		log.Fatal(err)
	}
	habitatImage = imaging.Resize(habitatImage, 36, 0, imaging.Box)

	foodImage, err := imaging.Open(foodPath)
	if err != nil {
		log.Fatal(err)
	}
	foodImage = imaging.Resize(foodImage, 36, 0, imaging.Box)

	nestingImage, err := imaging.Open(nestingPath)
	if err != nil {
		log.Fatal(err)
	}
	nestingImage = imaging.Resize(nestingImage, 36, 0, imaging.Box)

	behaviorImage, err := imaging.Open(behaviorPath)
	if err != nil {
		log.Fatal(err)
	}
	behaviorImage = imaging.Resize(behaviorImage, 36, 0, imaging.Box)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, habitatImage, image.Pt(0, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, foodImage, image.Pt(40, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, nestingImage, image.Pt(0, 40), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, behaviorImage, image.Pt(40, 40), 255)

	lifeHistoryImage = img.RgbaToGray(lifeHistoryImage)

	pngPath, bmpPath := config.LifeHistoryImagePaths()

	err = imaging.Save(lifeHistoryImage, pngPath)
	if err != nil {
		fmt.Println(err)
	}

	err = imaging.Save(lifeHistoryImage, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}
