package middleware

import "fmt"

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Home")

	result := middlewareOps(add)(2, 3)
	fmt.Println(result)

	result = middlewareOps(substract)(10, 3)
	fmt.Println(result)

	result = middlewareOps(multiply)(3, 3)
	fmt.Println(result)
}

func middlewareOps(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operación")
		return f(a, b)
	}
}
