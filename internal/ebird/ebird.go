package ebird

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bird struct {
	ScientificName       string `json:"sciName"`
	CommonName           string `json:"comName"`
	SpeciesCode          string `json:"speciesCode"`
	Category             string `json:"category"`
	Order                string `json:"order"`
	FamilyCode           string `json:"familyCode"`
	FamilyCommonName     string `json:"familyComName"`
	FamilyScientificName string `json:"familySciName"`
	ReportAs             string `json:"reportAs"`
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
