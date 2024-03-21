package main

import (
	"fmt"

	"github.com/jorojas/gofromscratch/variables"
)

func main() {
	var status, value = variables.ConvertToText(10)
	fmt.Println(value)
	fmt.Println(status)
}
