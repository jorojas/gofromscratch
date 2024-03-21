package main

import (
	"fmt"
	"runtime"

	"github.com/jorojas/gofromscratch/exercises"
	"github.com/jorojas/gofromscratch/variables"
)

func main() {
	var status, value = variables.ConvertToText(10)
	fmt.Println(value)
	fmt.Println(status)

	os := runtime.GOOS
	if os == "darwin" {
		fmt.Println("Hola soy Linux Mac - Darwin")
	} else {
	}

	if os1 := runtime.GOOS; os1 == "darwin" {
		fmt.Println("Hola soy Linux Mac - Darwin, un IF más óptimo")
	}

	switch os2 := runtime.GOOS; os2 {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s por defecto", os2)
	}

	num, msg := exercises.ConvertToInt("hola")

	fmt.Printf("%d %s\n", num, msg)
}
