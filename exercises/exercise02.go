package exercises

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ShowMultiplicationTable() {
	fmt.Println("Bienvenido a la tabla de multiplicar. Por favor ingresa un número del 1 al 10 : ")
	num, err := getNumberFromInput()

	for err != nil {
		num, err = getNumberFromInput()
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", num, i, (num * i))
	}
}

func getNumberFromInput() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("El dato ingresado es incorrecto. " + err.Error())
			return -1, err
		} else if num > 10 || num < 1 {
			fmt.Println("El dato ingresado es incorrecto. Debe ser de 1 al 10")
			return -1, errors.New(fmt.Sprint("Valor inválido ", num))
		} else {
			return num, nil
		}
	} else {
		panic("Could not open an input scanner")
	}
}
