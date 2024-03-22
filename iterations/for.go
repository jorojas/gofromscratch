package iterations

import "fmt"

func Iterate() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	for x := 0; x < 100; x += 5 {
		fmt.Println(x)
	}

}
