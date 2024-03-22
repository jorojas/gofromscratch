package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var legend string
var err error

func InputNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1 : ")

	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto. " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2 : ")

	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto. " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")

	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println(legend, num1*num2)
}
