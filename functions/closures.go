package functions

import "fmt"

func table(value int) func() int {
	number := value

	seq := 0

	return func() int {
		seq++
		return number * seq
	}
}

func CallClosure() {
	tableOf := 2
	MyTable := table(tableOf)

	for i := 1; i <= 10; i++ {
		fmt.Println(MyTable())
	}
}
