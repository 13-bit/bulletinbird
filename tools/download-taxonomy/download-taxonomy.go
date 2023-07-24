package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/13-bit/bulletinbird/birds"
	"github.com/13-bit/bulletinbird/util"
	"github.com/PuerkitoBio/goquery"
)

type BirdClass struct {
	TaxonCode string `json:"taxonCode"`
	Category  string `json:"category"`
	SciName   string `json:"sciName"`
	ComName   string `json:"comName"`
	SubTaxa   []struct {
		TaxonCode string `json:"taxonCode"`
		Category  string `json:"category"`
		SciName   string `json:"sciName"`
		ComName   string `json:"comName"`
		SubTaxa   []struct {
			TaxonCode string `json:"taxonCode"`
			Category  string `json:"category"`
			SciName   string `json:"sciName"`
			ComName   string `json:"comName"`
		} `json:"subTaxa"`
	} `json:"subTaxa"`
}

type BirdFamily struct {
	TaxonCode string `json:"taxonCode"`
	Category  string `json:"category"`
	SciName   string `json:"sciName"`
	ComName   string `json:"comName"`
	SubTaxa   []struct {
		TaxonCode string `json:"taxonCode"`
		Category  string `json:"category"`
		SciName   string `json:"sciName"`
		ComName   string `json:"comName"`
		SubTaxa   []struct {
			TaxonCode    string `json:"taxonCode"`
			Category     string `json:"category"`
			SciName      string `json:"sciName"`
			ComName      string `json:"comName"`
			IucnStatus   string `json:"iucnStatus"`
			IllusAssetID string `json:"illusAssetId"`
		} `json:"subTaxa,omitempty"`
	} `json:"subTaxa"`
}

func fetchBirdFamilies() []string {
	resp, err := http.Get("https://birdsoftheworld.org/bow/api/v1/taxonomy?depth=2&categoriesForCounts=species&showIucnStatusCounts=true&regionFilterId=US&locale=en")
	util.CheckError(err)

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	var bc BirdClass
	json.Unmarshal(bodyBytes, &bc)

	var families []string

	for _, orders := range bc.SubTaxa {
		for _, family := range orders.SubTaxa {
			families = append(families, family.TaxonCode)
		}
	}

	return families
}

func fetchIllustrationAssets() map[string]string {
	families := fetchBirdFamilies()

	var illustrationAssets = make(map[string]string)

	for _, family := range families {
		resp, err := http.Get(fmt.Sprintf("https://birdsoftheworld.org/bow/api/v1/taxonomy?depth=2&rootTaxonCode=%s&regionFilterId=US&locale=en", family))
		util.CheckError(err)

		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)

		var bf BirdFamily
		json.Unmarshal(bodyBytes, &bf)

		for _, genus := range bf.SubTaxa {
			for _, species := range genus.SubTaxa {
				illustrationAssets[species.SciName] = fmt.Sprintf("https://cdn.download.ams.birds.cornell.edu/api/v1/asset/%s/640", species.IllusAssetID)
			}
		}
	}

	return illustrationAssets
}

func downloadTaxonomyGuide() {
	log.Println("Downloading taxonomy...")

	illustrationAssets := fetchIllustrationAssets()

	var taxonomy []birds.Bird

	taxonRes, err := http.Get("https://www.allaboutbirds.org/guide/browse/taxonomy")
	util.CheckError(err)

	defer taxonRes.Body.Close()
	if taxonRes.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", taxonRes.StatusCode, taxonRes.Status)
	}

	guideDoc, err := goquery.NewDocumentFromReader(taxonRes.Body)
	util.CheckError(err)

	speciesCards := guideDoc.Find(".species-card")

	numSpecies := speciesCards.Length()
	speciesIndex := 0

	speciesCards.Each(func(i int, s *goquery.Selection) {
		commonName := s.Find(".species-info").Find("h4").Text()
		scientificName := s.Find(".species-info").Find("p").Text()
		guideRelativeUrl, _ := s.Find(".species-info").Find("h4").Find("a").Attr("href")
		guideUrl := fmt.Sprintf("https://www.allaboutbirds.org%s", guideRelativeUrl)

		speciesIndex += 1
		log.Printf("Downloading %s (%d of %d)... ", commonName, speciesIndex, numSpecies)

		guideRes, err := http.Get(guideUrl)
		util.CheckError(err)

		defer guideRes.Body.Close()
		if guideRes.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", guideRes.StatusCode, guideRes.Status)
		}

		speciesDoc, err := goquery.NewDocumentFromReader(guideRes.Body)
		util.CheckError(err)

		var lifeHistoryImageUrls []string

		speciesDoc.Find(".icon").Each(func(i int, s *goquery.Selection) {
			icon, _ := s.Find("img").Attr("src")

			lifeHistoryImageUrls = append(lifeHistoryImageUrls, fmt.Sprintf("https://www.allaboutbirds.org%s", icon))
		})

		bird := birds.Bird{}
		bird.CommonName = commonName
		bird.ScientificName = scientificName
		bird.GuideUrl = guideUrl
		bird.IllustrationUrl = illustrationAssets[scientificName]
		bird.LifeHistoryImageUrls = lifeHistoryImageUrls[0:5]

		stripUrl := func(url string) string {
			retStr := strings.ReplaceAll(url, "https://www.allaboutbirds.org/guide/images/icons/icon-", "")
			retStr = strings.ReplaceAll(retStr, ".png", "")

			return retStr
		}

		bird.Habitat = stripUrl(bird.LifeHistoryImageUrls[0])
		bird.Food = stripUrl(bird.LifeHistoryImageUrls[1])
		bird.Nesting = stripUrl(bird.LifeHistoryImageUrls[2])
		bird.Behavior = stripUrl(bird.LifeHistoryImageUrls[3])
		bird.Conservation = stripUrl(bird.LifeHistoryImageUrls[4])

		taxonomy = append(taxonomy, bird)

		log.Println("Done.")
	})

	log.Println("Saving taxonomy...")

	f, err := os.Create(taxonomyFilePath())
	util.CheckError(err)

	defer f.Close()

	taxonomyJson, _ := json.MarshalIndent(taxonomy, "", "  ")

	f.Write(taxonomyJson)

	log.Printf("%d birds saved to %s\n", len(taxonomy), taxonomyFilePath())
}

func taxonomyFilePath() string {
	moduleBaseDir := util.GoModDir()

	return fmt.Sprintf("%s/birds/resources/taxonomy.json", moduleBaseDir)
}

func main() {
	downloadTaxonomyGuide()
}
