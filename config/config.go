package config

import (
	"fmt"
	"os"
)

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
	staticPath string

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
	return fmt.Sprintf("%s/fonts/IBMPlexSans-Regular.ttf", configDir), fmt.Sprintf("%s/fonts/IBMPlexSans-Italic.ttf", configDir), fmt.Sprintf("%s/fonts/Pacifico-Regular.ttf", configDir)
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
