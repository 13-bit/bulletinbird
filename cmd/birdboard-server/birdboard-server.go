package main

import (
	"fmt"

	"github.com/13-bit/birdboard/internal/config"
	"github.com/13-bit/birdboard/internal/ebird"
)

type Configuration struct {
	EbirdToken string
}

func main() {
	fmt.Println("Welcome to Birdboard!")

	configuration := config.GetConfiguration()
	fmt.Println(configuration.EbirdToken)

	ebird.EbirdTest()
	ebird.DownloadTaxonomy()
}
