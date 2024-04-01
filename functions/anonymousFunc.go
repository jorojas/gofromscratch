package functions

import "fmt"

func Calculations() {
	add := func(value1 int, value2 int) int {
		return value1 + value2
	}

	fmt.Println(add(1, 2))
}
