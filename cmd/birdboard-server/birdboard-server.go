package main

import (
	"fmt"

	birds "github.com/13-bit/birdboard/internal/birds"
)

type Configuration struct {
	EbirdToken string
}

func main() {
	fmt.Println("Welcome to Birdboard!")

	// configuration := config.GetConfiguration()
	// fmt.Println(configuration.EbirdToken)

	// ebird.EbirdTest()
	// ebird.DownloadTaxonomy()
	birds.GetBirdList()
	birds.BirdOfTheDay()
}
