package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/13-bit/birdboard/internal/config"
	"github.com/bits-and-blooms/bloom/v3"
)

type Taxon struct {
	ScientificName string `json:"sciName"`
	CommonName     string `json:"comName"`
	SpeciesCode    string `json:"speciesCode"`
	Category       string `json:"category"`
	ReportAs       string `json:"reportAs"`
}

func ResetBotd() {
	fmt.Println("Resetting BOTD...")

	pastBotd := bloom.NewWithEstimates(13000, 0.01)

	fmt.Println("Saving BOTD...")

	f, err := os.Create(config.BotdFilePath())
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxonomyJson, err := json.Marshal(pastBotd)

	f.Write(taxonomyJson)
}

func main() {
	ResetBotd()
}
