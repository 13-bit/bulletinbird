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
