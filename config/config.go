package config

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/13-bit/bulletinbird/util"
)

//go:embed resources/*
var resourcesFS embed.FS

var configDir string

var botdFilePath,
	birdListFilePath,
	botdImageDownloadPath,
	habitatDownloadPath,
	foodDownloadPath,
	nestingDownloadPath,
	behaviorDownloadPath,
	conservationDownloadPath,
	inkyImagePath,
	magtagImagePath,
	magtagLifeHistoryImagePath,
	magTagQrCodeImagePath,
	staticPath,
	regularFontPath,
	italicFontPath,
	scriptFontPath string

func init() {
	homeDir, _ := os.UserHomeDir()
	configDir = fmt.Sprintf("%s/.bulletinbird", homeDir)

	botdFilePath = fmt.Sprintf("%s/botd.json", configDir)
	birdListFilePath = fmt.Sprintf("%s/birdlist.json", configDir)
	botdImageDownloadPath = fmt.Sprintf("%s/tmp/botd.png", configDir)
	habitatDownloadPath = fmt.Sprintf("%s/tmp/life-history-habitat.png", configDir)
	foodDownloadPath = fmt.Sprintf("%s/tmp/life-history-food.png", configDir)
	nestingDownloadPath = fmt.Sprintf("%s/tmp/life-history-nesting.png", configDir)
	behaviorDownloadPath = fmt.Sprintf("%s/tmp/life-history-behavior.png", configDir)
	conservationDownloadPath = fmt.Sprintf("%s/tmp/life-history-conservation.png", configDir)
	inkyImagePath = fmt.Sprintf("%s/static/inky/botd.png", configDir)
	magtagImagePath = fmt.Sprintf("%s/static/magtag/botd.bmp", configDir)
	magtagLifeHistoryImagePath = fmt.Sprintf("%s/static/magtag/life-history.bmp", configDir)
	magTagQrCodeImagePath = fmt.Sprintf("%s/static/magtag/qr.bmp", configDir)
	staticPath = fmt.Sprintf("%s/static", configDir)
	regularFontPath = fmt.Sprintf("%s/fonts/IBMPlexSans-Regular.ttf", configDir)
	italicFontPath = fmt.Sprintf("%s/fonts/IBMPlexSans-Italic.ttf", configDir)
	scriptFontPath = fmt.Sprintf("%s/fonts/Pacifico-Regular.ttf", configDir)

	if !checkConfigExists() {
		fmt.Printf("%s does not exist, creating...\n", configDir)
		createConfig()
	}
}

func checkConfigExists() bool {
	_, err := os.Stat(configDir)

	return !os.IsNotExist(err)
}

func createConfig() {
	if err := os.Mkdir(configDir, os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/static", configDir), os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/static/inky", configDir), os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/static/magtag", configDir), os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/tmp", configDir), os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/fonts", configDir), os.FileMode(0755)); err != nil {
		util.CheckError(err)
	}

	copyResourceFile("resources/IBMPlexSans-Regular.ttf", regularFontPath)
	copyResourceFile("resources/IBMPlexSans-Italic.ttf", italicFontPath)
	copyResourceFile("resources/Pacifico-Regular.ttf", scriptFontPath)
}

func copyResourceFile(readFilename string, writeFilename string) {
	log.Printf("Copying %s to %s...\n", readFilename, writeFilename)

	bytes, err := resourcesFS.ReadFile(readFilename)
	util.CheckError(err)

	err = os.WriteFile(writeFilename, bytes, os.FileMode(0755))
	util.CheckError(err)
}

func BotdFilePath() string {
	return botdFilePath
}

func BirdListFilePath() string {
	return birdListFilePath
}

func BotdImageDownloadPath() string {
	return botdImageDownloadPath
}

func QrCodeImageDownloadPath(size int) string {
	return fmt.Sprintf("%s/tmp/qr-%d.png", configDir, size)
}

func LifeHistoryImageDownloadPaths() (string, string, string, string, string) {
	return habitatDownloadPath,
		foodDownloadPath,
		nestingDownloadPath,
		behaviorDownloadPath,
		conservationDownloadPath
}

func MagTagQrCodeImagePath() string {
	return magTagQrCodeImagePath
}

func MagtagLifeHistoryImagePath() string {
	return magtagLifeHistoryImagePath
}

func StaticPath() string {
	return staticPath
}

func FontPaths() (string, string, string) {
	return regularFontPath, italicFontPath, scriptFontPath
}

func InkyImagePath() string {
	return inkyImagePath
}

func MagtagImagePath() string {
	return magtagImagePath
}

func InkyImageScript() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/inky/examples/7color/image.py", homeDir)
}
