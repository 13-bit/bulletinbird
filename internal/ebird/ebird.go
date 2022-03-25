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
	Name                 string   `json:"name"`
	Quantity             int      `json:"quantity"`
	ScientificName       string   `json:"sciName"`
	CommonName           string   `json:"comName"`
	SpeciesCode          string   `json:"speciesCode"`
	Category             string   `json:"category"`
	TaxonOrder           float64  `json:"taxonOrder"`
	BandingCodes         []string `json:"bandingCodes"`
	CommonNameCodes      []string `json:"comNameCodes"`
	ScientificNameCodes  []string `json:"sciNameCodes"`
	Order                string   `json:"order"`
	FamilyCode           string   `json:"familyCode"`
	FamilyCommonName     string   `json:"familyComName"`
	FamilyScientificName string   `json:"familySciName"`
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

	resp, err := http.Get("https://api.ebird.org/v2/ref/taxonomy/ebird?fmt=json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(bodyBytes)
	var taxa []Taxon

	err = json.Unmarshal(bodyBytes, &taxa)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(taxonomyFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxaJson, err := json.MarshalIndent(taxa, "", "  ")

	f.Write(taxaJson)
}
