package birds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/13-bit/birdboard/internal/config"
	"github.com/bits-and-blooms/bloom/v3"
)

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

func BirdOfTheDay() {
	homeDir, _ := os.UserHomeDir()
	pastFilePath := fmt.Sprintf("%s/.birdboard/past.json", homeDir)

	configuration := config.GetConfiguration()
	fmt.Println(configuration.EbirdToken)

	filter := bloom.NewWithEstimates(13000, 0.01)
	filter.Add([]byte("amecro"))

	fmt.Printf(fmt.Sprintf("%v", filter))

	f, err := os.Create(pastFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	filterJson, err := json.MarshalIndent(filter, "", "  ")

	f.Write(filterJson)
}
