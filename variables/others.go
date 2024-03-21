package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func PrintOtherVars() {
	Name = "Jonathan"
	Status = true
	Salary = 1200.30
	Date = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToText(num int) (bool, string) {

	text := strconv.Itoa(num)

	return true, text
}
