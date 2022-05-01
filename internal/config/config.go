package config

import (
	"fmt"
	"os"
)

func TaxonomyFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)
}

func BotdFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/botd.json", homeDir)
}

func BirdListFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/birdlist.json", homeDir)
}

func BotdImageDownloadPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/botd.png", homeDir)
}

func QrCodeImageDownloadPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/qr.png", homeDir)
}

func LifeHistoryImageDownloadPaths() (string, string, string, string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/life-history-habitat.png", homeDir),
		fmt.Sprintf("%s/.birdboard/life-history-food.png", homeDir),
		fmt.Sprintf("%s/.birdboard/life-history-nesting.png", homeDir),
		fmt.Sprintf("%s/.birdboard/life-history-behavior.png", homeDir),
		fmt.Sprintf("%s/.birdboard/life-history-conservation.png", homeDir)
}

func BotdImageFilePaths() (string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static/botd.png", homeDir), fmt.Sprintf("%s/.birdboard/static/botd.bmp", homeDir)
}

func QrCodeImageFilePaths() (string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static/qr.png", homeDir), fmt.Sprintf("%s/.birdboard/static/qr.bmp", homeDir)
}

func LifeHistoryImagePaths() (string, string) {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static/life-history.png", homeDir), fmt.Sprintf("%s/.birdboard/static/life-history.bmp", homeDir)
}

func LifeHistoryTemplateImagePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/life-history-template.png", homeDir)
}

func StaticPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static", homeDir)
}
