package main

import (
	"github.com/13-bit/birdboard/internal/birds"
)

func ResetBotd() {
	birds.RegenerateBirdList()
}

func main() {
	ResetBotd()
}
