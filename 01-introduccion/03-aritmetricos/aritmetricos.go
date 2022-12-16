package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Suma
	result := a + b
	fmt.Println("Suma: ", result)

	result = a - b
	fmt.Println("Resta", result)

	result = a * b
	fmt.Println("multiplicacion", result)

	result = a / b
	fmt.Println("Division", result)

	var div float64 = 3.0 / 2.0
	fmt.Println("resultado:", div)

	result = 2 % 3
	fmt.Println("Module:", result)
}
