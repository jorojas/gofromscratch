package main

import (
	"fmt"
	"runtime"

	exinterfaces "github.com/jorojas/gofromscratch/ex_interfaces"
	"github.com/jorojas/gofromscratch/exercises"
	"github.com/jorojas/gofromscratch/goroutines"
	"github.com/jorojas/gofromscratch/keyboard"
	"github.com/jorojas/gofromscratch/middleware"
	"github.com/jorojas/gofromscratch/models"
	"github.com/jorojas/gofromscratch/variables"
)

func mainOFF() {
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

	keyboard.InputNumbers()
}

func main() {
	//fmt.Println(exercises.ShowMultiplicationTable())

	//files.AddTable()

	//files.ReadFile()

	//functions.Calculations()

	//functions.CallClosure()

	//functions.Exponencia(2)

	//arraysslices.ShowArrays()

	//arraysslices.ShowSlice()

	//arraysslices.Capacity()

	//maps.ShowMaps()

	//users.CreateUser()

	//execInterfaceEx()

	//deferpanic.PanicExample()

	//execGoRoutines()

	//webserver.MyWebServer()

	middleware.MyMiddleware()

}

func execGoRoutines() {
	channel1 := make(chan bool)

	go goroutines.MySlowName("Jonathan Rojas", channel1)

	defer func() {
		<-channel1
	}()
}

func execInterfaceEx() {
	Pedro := new(models.Man)
	Nataly := new(models.Woman)

	exinterfaces.GetHumansBreathing(Pedro)
	exinterfaces.GetHumansBreathing(Nataly)
}
