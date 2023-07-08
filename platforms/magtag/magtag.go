package magtag

import (
	"embed"
	"fmt"
	"image"

	"github.com/13-bit/bulletinbird/config"
	"github.com/13-bit/bulletinbird/img"
	"github.com/13-bit/bulletinbird/util"
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
	bmpPath := config.MagtagImagePath()

	botdImage, err := imaging.Open(botdPath)
	util.CheckError(err)

	if botdImage.Bounds().Dx() > botdImage.Bounds().Dy() {
		botdImage = imaging.Resize(botdImage, 100, 0, imaging.Box)
	} else {
		botdImage = imaging.Resize(botdImage, 0, 80, imaging.Box)
	}

	botdImage = img.RgbaToGray(botdImage)

	err = imaging.Save(botdImage, bmpPath)
	util.CheckError(err)
}

func processQrCodeImage() {
	fmt.Println("Processing QR code...")

	qrPath := config.QrCodeImageDownloadPath(80)

	bmpPath := config.MagTagQrCodeImagePath()

	qrImage, err := imaging.Open(qrPath)
	util.CheckError(err)

	qrImage = img.RgbaToGray(qrImage)

	err = imaging.Save(qrImage, bmpPath)
	util.CheckError(err)
}

func processLifeHistoryImages() {
	fmt.Println("Processing life history images...")

	lifeHistoryFile, err := resourcesFS.Open("resources/life-history-template.png")
	util.CheckError(err)

	defer lifeHistoryFile.Close()

	lifeHistoryImage, _, err := image.Decode(lifeHistoryFile)
	util.CheckError(err)

	habitatPath, foodPath, nestingPath, behaviorPath, _ := config.LifeHistoryImageDownloadPaths()

	habitatImage, err := imaging.Open(habitatPath)
	util.CheckError(err)

	habitatImage = imaging.Resize(habitatImage, 36, 0, imaging.Box)

	foodImage, err := imaging.Open(foodPath)
	util.CheckError(err)

	foodImage = imaging.Resize(foodImage, 36, 0, imaging.Box)

	nestingImage, err := imaging.Open(nestingPath)
	util.CheckError(err)

	nestingImage = imaging.Resize(nestingImage, 36, 0, imaging.Box)

	behaviorImage, err := imaging.Open(behaviorPath)
	util.CheckError(err)

	behaviorImage = imaging.Resize(behaviorImage, 36, 0, imaging.Box)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, habitatImage, image.Pt(0, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, foodImage, image.Pt(40, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, nestingImage, image.Pt(0, 40), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, behaviorImage, image.Pt(40, 40), 255)

	lifeHistoryImage = img.RgbaToGray(lifeHistoryImage)

	bmpPath := config.MagtagLifeHistoryImagePath()

	err = imaging.Save(lifeHistoryImage, bmpPath)
	util.CheckError(err)
}
