package config

import (
	"fmt"
	"os"
)

func TaxonomyFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/taxonomy.json", homeDir)
}

func BotdFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/botd.json", homeDir)
}

func BirdListFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/birdlist.json", homeDir)
}

func BotdImageDownloadPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/botd.png", homeDir)
}

func QrCodeImageDownloadPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/qr.png", homeDir)
}

func LifeHistoryImageDownloadPaths() (string, string, string, string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/life-history-habitat.png", homeDir),
		fmt.Sprintf("%s/.bulletinbird/life-history-food.png", homeDir),
		fmt.Sprintf("%s/.bulletinbird/life-history-nesting.png", homeDir),
		fmt.Sprintf("%s/.bulletinbird/life-history-behavior.png", homeDir),
		fmt.Sprintf("%s/.bulletinbird/life-history-conservation.png", homeDir)
}

func BotdImageFilePath(format string) string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static/botd.%s", homeDir, format)
}

func QrCodeImageFilePath(format string) string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static/qr.%s", homeDir, format)
}

func LifeHistoryImagePath(format string) string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static/life-history.%s", homeDir, format)
}

func LifeHistoryTemplateImagePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/life-history-template.png", homeDir)
}

func StaticPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static", homeDir)
}

func InkyImagePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/inky/botd.png", homeDir)
}
