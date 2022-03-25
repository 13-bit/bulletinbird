package birds

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Bird struct {
	ScientificName       string `json:"sciName"`
	CommonName           string `json:"comName"`
	SpeciesCode          string `json:"speciesCode"`
	Category             string `json:"category"`
	Order                string `json:"order"`
	FamilyCode           string `json:"familyCode"`
	FamilyCommonName     string `json:"familyComName"`
	FamilyScientificName string `json:"familySciName"`
	ReportAs             string `json:"reportAs"`
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
