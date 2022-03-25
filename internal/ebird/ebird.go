package ebird

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Taxon struct {
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

func EbirdTest() {
	resp, err := http.Get("https://api.ebird.org/v2/ref/taxonomy/ebird?fmt=json&species=amecro")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func DownloadTaxonomy() {
	homeDir, _ := os.UserHomeDir()
	taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

	resp, err := http.Get("https://api.ebird.org/v2/ref/taxonomy/ebird?fmt=json&category=species")
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

	var filteredTaxonomy []Taxon

	for _, taxon := range taxonomy {
		if taxon.ReportAs == "" {
			filteredTaxonomy = append(filteredTaxonomy, taxon)
		}
	}

	f, err := os.Create(taxonomyFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxonomyJson, err := json.MarshalIndent(filteredTaxonomy, "", "  ")

	f.Write(taxonomyJson)
}
