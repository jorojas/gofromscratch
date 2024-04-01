package arraysslices

import "fmt"

var tableS []int = []int{2, 3, 4}

var array [10]int = [10]int{6, 33, 123, 30, 130, 32, 15, 76, 33, 23}

func ShowSlice() {
	fmt.Println(tableS)

	portion := array[3:] // slice created from vector from position 3
	portion2 := array[:5]
	portion3 := array[4:7]

	fmt.Println(portion)
	fmt.Println(portion2)
	fmt.Println(portion3)
}

func Capacity() {
	elements := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacity %d", len(elements), cap(elements))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Println()
	fmt.Printf("Largo %d, Capacity %d", len(nums), cap(nums))
	fmt.Println()
}
