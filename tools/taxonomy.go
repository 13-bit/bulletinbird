package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/13-bit/birdboard/internal/ebird"
)

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
	var taxonomy []ebird.Bird

	err = json.Unmarshal(bodyBytes, &taxonomy)
	if err != nil {
		log.Fatal(err)
	}

	var filteredTaxonomy []ebird.Bird

	for _, bird := range taxonomy {
		if bird.ReportAs == "" {
			filteredTaxonomy = append(filteredTaxonomy, bird)
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

func main() {
	DownloadTaxonomy()
}
