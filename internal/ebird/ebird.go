package ebird

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
