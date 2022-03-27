package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/PuerkitoBio/goquery"
)

// type Taxon struct {
// 	ScientificName string `json:"sciName"`
// 	CommonName     string `json:"comName"`
// 	SpeciesCode    string `json:"speciesCode"`
// 	Category       string `json:"category"`
// 	ReportAs       string `json:"reportAs"`
// }

// func DownloadTaxonomyEbird() {
// 	homeDir, _ := os.UserHomeDir()
// 	taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

// 	fmt.Println("Downloading taxonomy...")

// 	resp, err := http.Get("https://api.ebird.org/v2/ref/taxonomy/ebird?fmt=json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// bodyString := string(bodyBytes)
// 	var taxonomy []Taxon

// 	err = json.Unmarshal(bodyBytes, &taxonomy)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Filtering taxonomy...")

// 	var filteredTaxonomy []birds.Bird

// 	for _, taxon := range taxonomy {
// 		if taxon.ReportAs == "" {
// 			if taxon.Category == "species" {
// 				url := birds.EbirdUrl(taxon.SpeciesCode)

// 				if err == nil {
// 					bird := birds.Bird{}
// 					bird.ScientificName = taxon.ScientificName
// 					bird.CommonName = taxon.CommonName
// 					bird.SpeciesCode = taxon.SpeciesCode
// 					bird.EbirdUrl = url

// 					filteredTaxonomy = append(filteredTaxonomy, bird)
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("Saving taxonomy...")

// 	f, err := os.Create(taxonomyFilePath)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer f.Close()

// 	taxonomyJson, err := json.MarshalIndent(filteredTaxonomy, "", "  ")

// 	f.Write(taxonomyJson)

// 	fmt.Printf("%d birds saved to %s.\n", len(filteredTaxonomy), taxonomyFilePath)
// }

func DownloadTaxonomyGuide() {
	homeDir, _ := os.UserHomeDir()
	taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

	fmt.Println("Downloading taxonomy...")

	res, err := http.Get("https://www.allaboutbirds.org/guide/browse/taxonomy")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var taxonomy []birds.Bird

	doc.Find(".species-card").Each(func(i int, s *goquery.Selection) {
		commonName := s.Find(".species-info").Find("h4").Text()
		scientificName := s.Find(".species-info").Find("p").Text()
		guideRelativeUrl, _ := s.Find(".species-info").Find("h4").Find("a").Attr("href")
		guideUrl := fmt.Sprintf("https://www.allaboutbirds.org%s", guideRelativeUrl)

		bird := birds.Bird{}
		bird.CommonName = commonName
		bird.ScientificName = scientificName
		bird.GuideUrl = guideUrl

		taxonomy = append(taxonomy, bird)
	})

	fmt.Println("Saving taxonomy...")

	f, err := os.Create(taxonomyFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxonomyJson, err := json.MarshalIndent(taxonomy, "", "  ")

	f.Write(taxonomyJson)

	fmt.Printf("%d birds saved to %s.\n", len(taxonomy), taxonomyFilePath)
}

func main() {
	DownloadTaxonomyGuide()
}
