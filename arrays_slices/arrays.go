package arraysslices

import "fmt"

var table [10]int = [10]int{10, 0, 33}

func ShowArrays() {
	table[7] = 33
	table[2] = 54

	table2 := [10]string{"Pablo", "Juan"}

	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}
