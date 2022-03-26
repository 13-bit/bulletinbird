package botd

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"

	"github.com/13-bit/birdboard/internal/birds"
	"github.com/bits-and-blooms/bloom/v3"
)

var pastBotd = bloom.NewWithEstimates(13000, 0.01)
var once sync.Once

func GetPastBotd() *bloom.BloomFilter {
	once.Do(func() {
		homeDir, _ := os.UserHomeDir()
		botdFilePath := fmt.Sprintf("%s/.birdboard/botd.json", homeDir)

		// Load taxonomy from file
		f, err := os.Open(botdFilePath)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		err = json.NewDecoder(f).Decode(&pastBotd)
		if err != nil {
			log.Fatal(err)
		}
	})

	return pastBotd
}

// func EbirdTest() {
// 	resp, err := http.Get("https://api.ebird.org/v2/ref/taxonomy/ebird?fmt=json&species=amecro")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)
// }

func BirdOfTheDay() {
	botd := birds.Bird{}
	pastBotd := GetPastBotd()

	for {
		idx, e := rand.Int(rand.Reader, big.NewInt(1000))
		if e != nil {
			fmt.Print(e)
		}

		botd = birds.GetBirdList()[idx.Int64()]

		if !pastBotd.Test([]byte(botd.SpeciesCode)) {
			pastBotd.Add([]byte(botd.SpeciesCode))
			break
		}
	}

	SaveBotd()

	fmt.Printf("BOTD: %+v\n", botd)
}

func SaveBotd() {
	homeDir, _ := os.UserHomeDir()
	botdFilePath := fmt.Sprintf("%s/.birdboard/botd.json", homeDir)

	pastBotd := GetPastBotd()

	f, err := os.Create(botdFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	taxonomyJson, err := json.Marshal(pastBotd)

	f.Write(taxonomyJson)
}
