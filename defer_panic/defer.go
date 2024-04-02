package deferpanic

import (
	"fmt"
	"log"
)

func ReviewDefer() {
	fmt.Println("This is a first message")
	defer fmt.Println("This is a final message")

	fmt.Println("This is a second message")

	defer fmt.Println("Another message")
}

func PanicExample() {

	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("An error has ocurred and a panic was generated -> %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("Value 1 was found")
	}

}
