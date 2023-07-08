package inky

import (
	"embed"
	"image"
	"image/color"
	"log"
	"os/exec"
	"runtime"

	"github.com/13-bit/bulletinbird/birds"
	"github.com/13-bit/bulletinbird/config"
	"github.com/13-bit/bulletinbird/util"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

//go:embed resources/*
var resourcesFS embed.FS

const (
	inkyWidth         = 640
	inkyHeight        = 400
	botdWidth         = 420
	lifeHistoryWidth  = 216
	lifeHistoryHeight = 216
	iconSize          = 64
	qrSize            = 144
)

func Refresh() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("python3", config.InkyImageScript(), config.InkyImagePath())
		if err := cmd.Run(); err != nil {
			util.CheckError(err)
		}
	}
}

func GenerateInkyImages(botd birds.Bird) {
	log.Println("Generating images for Inky...")

	inkyImage := imaging.New(inkyWidth, inkyHeight, color.NRGBA{255, 255, 255, 255})

	botdImage, xOffset, yOffset := genBotdImage()
	botdX := inkyWidth - botdImage.Bounds().Dx() - xOffset

	inkyImage = imaging.Overlay(inkyImage, botdImage, image.Pt(botdX, yOffset), 255)

	textImage := genBotdText(botd)

	inkyImage = imaging.Overlay(inkyImage, textImage, image.Pt(0, 0), 255)

	lifeHistoryImage := genLifeHistoryImage()
	lifeHistoryY := inkyHeight - lifeHistoryHeight

	inkyImage = imaging.Overlay(inkyImage, lifeHistoryImage, image.Pt(0, lifeHistoryY), 255)

	err := imaging.Save(inkyImage, config.InkyImagePath())
	util.CheckError(err)
}

func genBotdImage() (image.Image, int, int) {
	log.Println("Generating BOTD image for Inky...")

	xOffset := 0
	yOffset := 0

	botdImage, err := imaging.Open(config.BotdImageDownloadPath())
	util.CheckError(err)

	if botdImage.Bounds().Dx() > botdImage.Bounds().Dy() {
		botdImage = imaging.Resize(botdImage, botdWidth, 0, imaging.Box)
		yOffset = (inkyHeight - botdImage.Bounds().Dy()) / 2
	} else {
		botdImage = imaging.Resize(botdImage, 0, inkyHeight, imaging.Box)
		xOffset = (botdWidth - botdImage.Bounds().Dx()) / 2
	}

	botdImage = imaging.AdjustSaturation(botdImage, 64)

	return botdImage, xOffset, yOffset
}

func genLifeHistoryImage() image.Image {
	log.Println("Generating life history images for Inky...")

	habitatPath, foodPath, nestingPath, behaviorPath, conservationPath := config.LifeHistoryImageDownloadPaths()

	habitatImage, err := imaging.Open(habitatPath)
	util.CheckError(err)

	habitatImage = imaging.Resize(habitatImage, iconSize, 0, imaging.Box)

	habitatIconFile, err := resourcesFS.Open("resources/icon-habitat.png")
	util.CheckError(err)

	defer habitatIconFile.Close()

	habitatIcon, _, err := image.Decode(habitatIconFile)
	util.CheckError(err)

	foodImage, err := imaging.Open(foodPath)
	util.CheckError(err)

	foodImage = imaging.Resize(foodImage, iconSize, 0, imaging.Box)

	foodIconFile, err := resourcesFS.Open("resources/icon-food.png")
	util.CheckError(err)

	defer foodIconFile.Close()

	foodIcon, _, err := image.Decode(foodIconFile)
	util.CheckError(err)

	nestingImage, err := imaging.Open(nestingPath)
	util.CheckError(err)

	nestingImage = imaging.Resize(nestingImage, iconSize, 0, imaging.Box)

	nestingIconFile, err := resourcesFS.Open("resources/icon-nesting.png")
	util.CheckError(err)

	defer nestingIconFile.Close()

	nestingIcon, _, err := image.Decode(nestingIconFile)
	util.CheckError(err)

	behaviorImage, err := imaging.Open(behaviorPath)
	util.CheckError(err)

	behaviorImage = imaging.Resize(behaviorImage, iconSize, 0, imaging.Box)

	behaviorIconFile, err := resourcesFS.Open("resources/icon-behavior.png")
	util.CheckError(err)

	defer behaviorIconFile.Close()

	behaviorIcon, _, err := image.Decode(behaviorIconFile)
	util.CheckError(err)

	conservationImage, err := imaging.Open(conservationPath)
	util.CheckError(err)

	conservationImage = imaging.Resize(conservationImage, iconSize, 0, imaging.Box)

	maskFile, err := resourcesFS.Open("resources/icon-mask.png")
	util.CheckError(err)

	defer behaviorIconFile.Close()

	iconMask, _, err := image.Decode(maskFile)
	util.CheckError(err)

	qrImage, err := imaging.Open(config.QrCodeImageDownloadPath(qrSize))
	util.CheckError(err)

	// qrImage = imaging.CropAnchor(qrImage, qrSize, qrSize, imaging.Center)

	lifeHistoryImage := imaging.New(lifeHistoryWidth, lifeHistoryHeight, color.NRGBA{255, 255, 255, 255})

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, qrImage, image.Pt(8, 64), 255)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, habitatIcon, image.Pt(8, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, habitatImage, image.Pt(8, 0), 255)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, foodIcon, image.Pt(80, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, foodImage, image.Pt(80, 0), 255)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, nestingIcon, image.Pt(152, 72), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, nestingImage, image.Pt(152, 72), 255)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, behaviorIcon, image.Pt(152, 144), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, behaviorImage, image.Pt(152, 144), 255)

	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, conservationImage, image.Pt(152, 0), 255)
	lifeHistoryImage = imaging.Overlay(lifeHistoryImage, iconMask, image.Pt(152, 0), 255)

	return lifeHistoryImage
}

func genBotdText(botd birds.Bird) image.Image {
	commonNameFont, scientificNameFont, titleFont := config.FontPaths()

	dc := gg.NewContext(220, 240)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace(titleFont, 28); err != nil {
		panic(err)
	}

	// dc.DrawStringWrapped("Bird of the Day", 16, 132, 16, 132, 200, 1.5, gg.AlignLeft)
	dc.DrawString("Bird of the Day", 16, 44)

	if err := dc.LoadFontFace(commonNameFont, 20); err != nil {
		panic(err)
	}

	lines := dc.WordWrap(botd.CommonName, 200)

	dc.DrawStringWrapped(botd.CommonName, 16, 80, 0, 0, 200, 1.25, gg.AlignLeft)
	// dc.DrawString(botd.CommonName, 16, 132)

	_, sh := dc.MeasureString(botd.CommonName)

	if err := dc.LoadFontFace(scientificNameFont, 16); err != nil {
		panic(err)
	}

	textOffset := (sh + 8) * float64(len(lines))

	dc.DrawStringWrapped(botd.ScientificName, 16, 80+textOffset, 0, 0, 200, 1.25, gg.AlignLeft)
	// dc.DrawString(botd.ScientificName, 16, 148)

	return dc.Image()
}
