package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	birds "github.com/13-bit/birdboard/internal/birds"
)

type Taxon struct {
	ScientificName string `json:"sciName"`
	CommonName     string `json:"comName"`
	SpeciesCode    string `json:"speciesCode"`
	Category       string `json:"category"`
	ReportAs       string `json:"reportAs"`
}

func DownloadTaxonomy() {
	homeDir, _ := os.UserHomeDir()
	taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

	fmt.Println("Downloading taxonomy...")

	resp, err := http.Get("https://api.ebird.org/v2/ref/taxonomy/ebird?fmt=json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(bodyBytes)
	var taxonomy []Taxon

	err = json.Unmarshal(bodyBytes, &taxonomy)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Filtering taxonomy...")

	var filteredTaxonomy []birds.Bird

	for _, taxon := range taxonomy {
		if taxon.ReportAs == "" {
			if taxon.Category == "species" {
				bird := birds.Bird{}
				bird.ScientificName = taxon.ScientificName
				bird.CommonName = taxon.CommonName
				bird.SpeciesCode = taxon.SpeciesCode

				filteredTaxonomy = append(filteredTaxonomy, bird)
			}
		}
	}

	fmt.Println("Saving taxonomy...")

	f, err := os.Create(taxonomyFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxonomyJson, err := json.MarshalIndent(filteredTaxonomy, "", "  ")

	f.Write(taxonomyJson)

	fmt.Printf("%d birds saved to %s.\n", len(filteredTaxonomy), taxonomyFilePath)
}

func main() {
	DownloadTaxonomy()
}
