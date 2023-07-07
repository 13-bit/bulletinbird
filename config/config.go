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
	conservationDownloadPath string

func init() {
	homeDir, _ := os.UserHomeDir()
	configDir = fmt.Sprintf("%s/.bulletinbird", homeDir)

	botdFilePath = fmt.Sprintf("%s/botd.json", configDir)
	birdListFilePath = fmt.Sprintf("%s/birdlist.json", configDir)
	botdImageDownloadPath = fmt.Sprintf("%s/botd.png", configDir)
	habitatDownloadPath = fmt.Sprintf("%s/life-history-habitat.png", configDir)
	foodDownloadPath = fmt.Sprintf("%s/life-history-food.png", configDir)
	nestingDownloadPath = fmt.Sprintf("%s/life-history-nesting.png", configDir)
	behaviorDownloadPath = fmt.Sprintf("%s/life-history-behavior.png", configDir)
	conservationDownloadPath = fmt.Sprintf("%s/life-history-conservation.png", configDir)
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
	return fmt.Sprintf("%s/qr-%d.png", configDir, size)
}

func LifeHistoryImageDownloadPaths() (string, string, string, string, string) {
	return habitatDownloadPath,
		foodDownloadPath,
		nestingDownloadPath,
		behaviorDownloadPath,
		conservationDownloadPath
}

func BotdImageFilePath(format string) string {
	return fmt.Sprintf("%s/static/botd.%s", configDir, format)
}

func QrCodeImageFilePath(format string) string {
	return fmt.Sprintf("%s/static/qr.%s", configDir, format)
}

func LifeHistoryImagePath(format string) string {
	return fmt.Sprintf("%s/static/life-history.%s", configDir, format)
}

func LifeHistoryTemplateImagePath() string {
	return fmt.Sprintf("%s/life-history-template.png", configDir)
}

func StaticPath() string {
	return fmt.Sprintf("%s/static", configDir)
}

func FontPaths() (string, string, string) {
	return fmt.Sprintf("%s/fonts/IBMPlexSans-Regular.ttf", configDir), fmt.Sprintf("%s/fonts/IBMPlexSans-Italic.ttf", configDir), fmt.Sprintf("%s/fonts/Pacifico-Regular.ttf", configDir)
}

func InkyImagePath() string {
	return fmt.Sprintf("%s/inky/botd.png", configDir)
}

func InkyImageScript() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/inky/examples/7color/image.py", homeDir)
}
