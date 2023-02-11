package birds

import (
	"container/list"
	cr "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	mr "math/rand"
	"os"
	"time"

	"github.com/13-bit/bulletinbird/config"
)

type Bird struct {
	ScientificName       string   `json:"sciName"`
	CommonName           string   `json:"comName"`
	GuideUrl             string   `json:"guideUrl"`
	IllustrationUrl      string   `json:"illustrationUrl"`
	LifeHistoryImageUrls []string `json:"lifeHistoryImageUrls"`
}

func GetTaxonomy() []Bird {
	f, err := os.Open(config.TaxonomyFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var taxonomy []Bird

	err = json.NewDecoder(f).Decode(&taxonomy)
	if err != nil {
		log.Fatal(err)
	}

	return taxonomy
}

func RegenerateBirdList() {
	taxonomy := GetTaxonomy()
	birdList := list.New()

	seed, e := cr.Int(cr.Reader, big.NewInt(time.Now().Unix()))
	if e != nil {
		fmt.Print(e)
	}

	mr.Seed(seed.Int64())
	randomIndices := mr.Perm(len(taxonomy))

	for _, idx := range randomIndices {
		birdList.PushBack(taxonomy[idx])
	}

	SaveBirdList(birdList)
}

func GetBirdList() *list.List {
	f, err := os.Open(config.BirdListFilePath())
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var birds []Bird

	err = json.NewDecoder(f).Decode(&birds)
	if err != nil {
		log.Fatal(err)
	}

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
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		birds := []Bird{}

		for e := birdList.Front(); e != nil; e = e.Next() {
			birds = append(birds, e.Value.(Bird))
		}

		birdListJson, _ := json.MarshalIndent(birds, "", "  ")

		f.Write(birdListJson)

		fmt.Printf("%d birds saved to %s\n", len(birds), config.BirdListFilePath())
	}
}
