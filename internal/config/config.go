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

func BotdImageFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static/botd.png", homeDir)
}

func QrCodeFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static/qr.png", homeDir)
}

func StaticPath() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.birdboard/static", homeDir)
}
