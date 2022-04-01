package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/13-bit/birdboard/internal/config"
	"github.com/PuerkitoBio/goquery"
)

func DownloadTaxonomyGuide() {
	fmt.Println("Downloading taxonomy...")

	var taxonomy []birds.Bird

	doc, err := goquery.NewDocument("https://www.allaboutbirds.org/guide/browse/taxonomy")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".species-card").Each(func(i int, s *goquery.Selection) {
		commonName := s.Find(".species-info").Find("h4").Text()
		scientificName := s.Find(".species-info").Find("p").Text()
		guideRelativeUrl, _ := s.Find(".species-info").Find("h4").Find("a").Attr("href")
		guideUrl := fmt.Sprintf("https://www.allaboutbirds.org%s", guideRelativeUrl)
		imgUrl, _ := s.Find(".species-image").Find("a").Find("img").Attr("src")

		if imgUrl == "" {
			imgUrl, _ = s.Find(".species-image").Find("a").Find("img").Attr("pre-src")
		}

		bird := birds.Bird{}
		bird.CommonName = commonName
		bird.ScientificName = scientificName
		bird.GuideUrl = guideUrl
		bird.ImgUrl = imgUrl

		taxonomy = append(taxonomy, bird)
	})

	fmt.Println("Saving taxonomy...")

	f, err := os.Create(config.TaxonomyFilePath())
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxonomyJson, err := json.MarshalIndent(taxonomy, "", "  ")

	f.Write(taxonomyJson)

	fmt.Printf("%d birds saved to %s\n", len(taxonomy), config.TaxonomyFilePath())
}

func main() {
	DownloadTaxonomyGuide()
}
