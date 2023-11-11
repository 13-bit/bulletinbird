package birds

import (
	"container/list"
	cr "crypto/rand"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	mr "math/rand"
	"os"
	"time"

	"github.com/13-bit/bulletinbird/config"
	"github.com/13-bit/bulletinbird/util"
)

//go:embed resources/*
var resourcesFS embed.FS

type Bird struct {
	ScientificName       string   `json:"sciName"`
	CommonName           string   `json:"comName"`
	GuideUrl             string   `json:"guideUrl"`
	IllustrationUrl      string   `json:"illustrationUrl"`
	IllustrationImage    string   `json:"illustrationImage"`
	LifeHistoryImageUrls []string `json:"lifeHistoryImageUrls"`
	LifeHistoryImages    []string `json:"lifeHistoryImages"`
	Habitat              string   `json:"habitat"`
	Food                 string   `json:"food"`
	Nesting              string   `json:"nesting"`
	Behavior             string   `json:"behavior"`
	Conservation         string   `json:"conservation"`
}

func GetTaxonomy() []Bird {
	f, err := resourcesFS.Open("resources/taxonomy.json")
	util.CheckError(err)

	defer f.Close()

	var taxonomy []Bird

	err = json.NewDecoder(f).Decode(&taxonomy)
	util.CheckError(err)

	for _, b := range taxonomy {
		fmt.Println(b.CommonName)
		fmt.Printf("habitat: %s\n", b.LifeHistoryImageUrls[0])
		fmt.Printf("diet: %s\n", b.LifeHistoryImageUrls[1])
		fmt.Printf("nest: %s\n", b.LifeHistoryImageUrls[2])
		fmt.Printf("behavior: %s\n", b.LifeHistoryImageUrls[3])
		fmt.Printf("conservation: %s\n", b.LifeHistoryImageUrls[4])
		fmt.Println()
	}

	return taxonomy
}

func RegenerateBirdList() {
	taxonomy := GetTaxonomy()
	birdList := list.New()

	seed, err := cr.Int(cr.Reader, big.NewInt(time.Now().Unix()))
	util.CheckError(err)

	mr.Seed(seed.Int64())
	randomIndices := mr.Perm(len(taxonomy))

	for _, idx := range randomIndices {
		birdList.PushBack(taxonomy[idx])
	}

	SaveBirdList(birdList)
}

func GetBirdList() *list.List {
	f, err := os.Open(config.BirdListFilePath())
	util.CheckError(err)

	defer f.Close()

	var birds []Bird

	err = json.NewDecoder(f).Decode(&birds)
	util.CheckError(err)

	birdList := list.New()

	for _, bird := range birds {
		birdList.PushBack(bird)
	}

	return birdList
}

func SaveBirdList(birdList *list.List) {
	if birdList.Len() == 0 {
		RegenerateBirdList()
	} else {
		f, err := os.Create(config.BirdListFilePath())
		util.CheckError(err)

		defer f.Close()

		birds := []Bird{}

		for e := birdList.Front(); e != nil; e = e.Next() {
			birds = append(birds, e.Value.(Bird))
		}

		birdListJson, _ := json.MarshalIndent(birds, "", "  ")

		f.Write(birdListJson)

		log.Printf("%d birds saved to %s\n", len(birds), config.BirdListFilePath())
	}
}
