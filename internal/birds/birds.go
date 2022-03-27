package birds

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Bird struct {
	ScientificName string `json:"sciName"`
	CommonName     string `json:"comName"`
	GuideUrl       string `json:"guideUrl"`
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

func SaveBirdList() {
	homeDir, _ := os.UserHomeDir()
	taxonomyFilePath := fmt.Sprintf("%s/.birdboard/taxonomy.json", homeDir)

	f, err := os.Create(taxonomyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	taxonomyJson, err := json.MarshalIndent(GetBirdList(), "", "  ")

	f.Write(taxonomyJson)

	fmt.Printf("%d birds saved to %s.\n", len(GetBirdList()), taxonomyFilePath)
}

// func EbirdUrl(speciesCode string) (string, error) {
// 	url := fmt.Sprintf("https://www.allaboutbirds.org/guide/%s", speciesCode)

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	client := new(http.Client)
// 	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
// 		return errors.New("Redirect")
// 	}

// 	_, err = client.Do(req)
// 	if err != nil {
// 		fmt.Printf("-")
// 		return "", errors.New(fmt.Sprintf("No guide page found for sepcies code %s", speciesCode))
// 	}

// 	fmt.Printf("+")

// 	return url, nil
// }

func EbirdUrl(speciesCode string) string {
	url := fmt.Sprintf("https://www.allaboutbirds.org/guide/%s", speciesCode)

	return url
}
