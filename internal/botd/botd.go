package botd

import (
	"fmt"

	"github.com/13-bit/birdboard/internal/birds"
)

func BirdOfTheDay() {
	birdList := birds.GetBirdList()

	botd := birdList.Front()

	birdList.Remove(botd)
	birds.SaveBirdList(birdList)

	fmt.Println("BOTD:", botd.Value)
}
