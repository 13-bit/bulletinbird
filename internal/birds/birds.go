package birds

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Bird struct {
	ScientificName string `json:"sciName"`
	CommonName     string `json:"comName"`
	SpeciesCode    string `json:"speciesCode"`
}

var birdList []Bird
var once sync.Once

func GetBirdList() []Bird {
	once.Do(func() {
		homeDir, _ := os.UserHomeDir()
		taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

		// Load taxonomy from file
		f, err := os.Open(taxonomyFilePath)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		err = json.NewDecoder(f).Decode(&birdList)
		if err != nil {
			log.Fatal(err)
		}
	})

	return birdList
}

func SaveBirdList() {
	homeDir, _ := os.UserHomeDir()
	taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

	f, err := os.Create(taxonomyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	taxonomyJson, err := json.MarshalIndent(GetBirdList(), "", "  ")

	f.Write(taxonomyJson)

	fmt.Printf("%d birds saved to %s.\n", len(GetBirdList()), taxonomyFilePath)
}
