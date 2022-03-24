package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	EbirdToken string
}

var configuration Configuration
var once sync.Once

func GetConfiguration() Configuration {
	once.Do(func() {
		homeDir, _ := os.UserHomeDir()
		configFilePath := fmt.Sprintf("%s/.birdboard/config.json", homeDir)

		configuration = Configuration{}
		err := gonfig.GetConf(configFilePath, &configuration)
		if err != nil {
			panic(err)
		}
	})

	return configuration
}
