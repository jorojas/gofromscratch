package functions

import "fmt"

func Exponencia(value int) {
	if value > 1000000 {
		return
	}

	fmt.Println(value)
	Exponencia(value * 2)
}
