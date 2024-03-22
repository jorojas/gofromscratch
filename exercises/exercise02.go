package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowMultiplicationTable() string {
	fmt.Println("Bienvenido a la tabla de multiplicar. Por favor ingresa un número del 1 al 10 : ")
	num := getNumberFromInput()
	var text string

	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", num, i, (num * i))
	}

	return text
}

func getNumberFromInput() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("El dato ingresado es incorrecto. " + err.Error())
				continue
			} else if num > 10 || num < 1 {
				fmt.Println("El dato ingresado es incorrecto. Debe ser de 1 al 10")
				continue
			} else {
				return num
			}
		} else {
			panic("Could not open an input scanner")
		}
	}
}
