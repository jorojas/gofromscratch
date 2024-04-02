package exinterfaces

import (
	"fmt"

	"github.com/jorojas/gofromscratch/interfaces"
)

func GetHumansBreathing(hu interfaces.Human) {
	hu.Breath()

	fmt.Printf("I am a %s, and I'm breathing.", hu.GetGender())
	fmt.Println()
}
