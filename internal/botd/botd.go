package botd

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"
	"time"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/13-bit/birdboard/internal/config"
	"github.com/13-bit/birdboard/internal/img"
	"github.com/cavaliergopher/grab/v3"
	"github.com/disintegration/imaging"
	qrcode "github.com/skip2/go-qrcode"
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
	downloadBotdImage(botd)
	processBotdImage()
	generateQrCode(botd)

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

func downloadBotdImage(botd BirdOfTheDay) {
	fmt.Println("downloading image...")

	botdPath := config.BotdImageDownloadPath()

	err := os.Remove(botdPath)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := grab.Get(botdPath, botd.Bird.ImgUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)
}

func processBotdImage() {
	fmt.Println("processing image...")

	botdPath := config.BotdImageDownloadPath()
	pngPath, bmpPath := config.BotdImageFilePaths()

	botdFile, err := os.Open(botdPath)
	if err != nil {
		log.Fatal(err)
	}
	defer botdFile.Close()

	botdImage, _, err := image.Decode(botdFile)
	if err != nil {
		fmt.Println(err)
	}

	botdImageResized := imaging.Resize(botdImage, 100, 0, imaging.Box)
	botdImageGrayscale := img.RgbaToGray(botdImageResized)

	// palette := []color.Color{
	// 	color.Gray16{0},
	// 	color.Gray16{0x400f},
	// 	color.Gray16{0x7fff},
	// 	color.Gray16{0xffff},
	// }

	// d := dither.NewDitherer(palette)
	// d.Matrix = dither.Atkinson

	// _ = d.Dither(botdImageGrayscale)

	err = imaging.Save(botdImageGrayscale, pngPath)
	if err != nil {
		fmt.Println(err)
	}

	err = imaging.Save(botdImageGrayscale, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}

func generateQrCode(botd BirdOfTheDay) {
	pngPath, bmpPath := config.QrCodeFilePaths()

	err := qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Low, 40, pngPath)
	if err != nil {
		fmt.Println(err)
	}

	err = qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Low, 40, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}
