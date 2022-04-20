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
	Bird          birds.Bird    `json:"bird"`
	LastUpdated   time.Time     `json:"lastUpdated"`
	UntilTomorrow time.Duration `json:"untilTomorrow"`
}

func nextBirdOfTheDay() BirdOfTheDay {
	birdList := birds.GetBirdList()

	front := birdList.Front()
	bird := front.Value.(birds.Bird)

	birdList.Remove(front)
	birds.SaveBirdList(birdList)

	botd := BirdOfTheDay{
		Bird:          bird,
		LastUpdated:   time.Now(),
		UntilTomorrow: 0,
	}

	SaveBotd(botd)
	downloadBotdImage(botd)
	processBotdImage()
	generateQrCode(botd)
	processQrCodeImage()

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

func GetBirdOfTheDay() BirdOfTheDay {
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

	botd.UntilTomorrow = timeUntilTomorrow()

	fmt.Printf("%+v\n", botd.Bird)

	return botd
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

	botdImage = imaging.Resize(botdImage, 100, 0, imaging.Box)
	botdImage = img.RgbaToGray(botdImage)

	// palette := []color.Color{
	// 	color.Gray16{0},
	// 	color.Gray16{0x400f},
	// 	color.Gray16{0x7fff},
	// 	color.Gray16{0xffff},
	// }

	// d := dither.NewDitherer(palette)
	// d.Matrix = dither.Atkinson

	// botdImage = d.Dither(botdImage)

	err = imaging.Save(botdImage, pngPath)
	if err != nil {
		fmt.Println(err)
	}

	err = imaging.Save(botdImage, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}

func generateQrCode(botd BirdOfTheDay) {
	qrPath := config.QrCodeImageDownloadPath()

	err := qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Medium, 80, qrPath)
	if err != nil {
		fmt.Println(err)
	}

	// err = qrcode.WriteFile(botd.Bird.GuideUrl, qrcode.Low, 60, bmpPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func processQrCodeImage() {
	fmt.Println("processing image...")

	qrPath := config.QrCodeImageDownloadPath()

	pngPath, bmpPath := config.QrCodeImageFilePaths()

	qrFile, err := os.Open(qrPath)
	if err != nil {
		log.Fatal(err)
	}
	defer qrFile.Close()

	qrImage, _, err := image.Decode(qrFile)
	if err != nil {
		fmt.Println(err)
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

func tomorrow() time.Time {
	loc, _ := time.LoadLocation("America/Chicago")
	year, month, day := time.Now().In(loc).Date()
	tomorrowMorning := time.Date(year, month, day+1, 5, 0, 0, 0, loc)

	return tomorrowMorning
}

func timeUntilTomorrow() time.Duration {
	tomorrowMorning := tomorrow()

	return time.Duration(tomorrowMorning.Sub(time.Now()).Seconds())
}
