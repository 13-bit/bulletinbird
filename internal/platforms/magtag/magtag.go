package magtag

import (
	"embed"
	"fmt"
	"image"
	"log"

	"github.com/13-bit/bulletinbird-server/internal/config"
	"github.com/13-bit/bulletinbird-server/internal/img"
	"github.com/disintegration/imaging"
)

//go:embed resources/*
var resourcesFS embed.FS

func GenerateMagtagImages() {
	processBotdImage()
	processQrCodeImage()
	processLifeHistoryImages()
}

func processBotdImage() {
	fmt.Println("Processing image...")

	botdPath := config.BotdImageDownloadPath()
	bmpPath := config.BotdImageFilePath("bmp")

	botdImage, err := imaging.Open(botdPath)
	if err != nil {
		log.Fatal(err)
	}

	if botdImage.Bounds().Dx() > botdImage.Bounds().Dy() {
		botdImage = imaging.Resize(botdImage, 100, 0, imaging.Box)
	} else {
		botdImage = imaging.Resize(botdImage, 0, 80, imaging.Box)
	}

	botdImage = img.RgbaToGray(botdImage)

	err = imaging.Save(botdImage, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}

func processQrCodeImage() {
	fmt.Println("Processing QR code...")

	qrPath := config.QrCodeImageDownloadPath()

	bmpPath := config.QrCodeImageFilePath("bmp")

	qrImage, err := imaging.Open(qrPath)
	if err != nil {
		log.Fatal(err)
	}

	qrImage = img.RgbaToGray(qrImage)

	err = imaging.Save(qrImage, bmpPath)
	if err != nil {
		fmt.Println(err)
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

	bmpPath := config.LifeHistoryImagePath("bmp")

	err = imaging.Save(lifeHistoryImage, bmpPath)
	if err != nil {
		fmt.Println(err)
	}
}
