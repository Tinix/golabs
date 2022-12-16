package main

import "fmt"

func cociente(a, b int) int {
	return a / b
}

func resto(a, b int) int {
	return a % b
}

func main() {
	var a, b, c, result int

	// Enter number one
	fmt.Println("Enter number one")
	fmt.Scanln(&a)

	// Enter number two
	fmt.Println("Enter number two")
	fmt.Scanln(&b)

	// call fucntion cociente
	c = cociente(a, b)

	result = resto(a, b)

	// Output data
	fmt.Println("Cociente", c)
	fmt.Println("Resto", result)

}
