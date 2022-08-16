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

func BotdImageFilePaths() (string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static/botd.png", homeDir), fmt.Sprintf("%s/.bulletinbird/static/botd.bmp", homeDir)
}

func QrCodeImageFilePaths() (string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static/qr.png", homeDir), fmt.Sprintf("%s/.bulletinbird/static/qr.bmp", homeDir)
}

func LifeHistoryImagePaths() (string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static/life-history.png", homeDir), fmt.Sprintf("%s/.bulletinbird/static/life-history.bmp", homeDir)
}

func LifeHistoryTemplateImagePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/life-history-template.png", homeDir)
}

func StaticPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.bulletinbird/static", homeDir)
}
