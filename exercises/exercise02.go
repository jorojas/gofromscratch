package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowMultiplicationTable() {
	fmt.Println("Bienvenido a la tabla de multiplicar. Por favor ingresa un número del 1 al 10 : ")
	num := getNumberFromInput()

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", num, i, (num * i))
	}
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
