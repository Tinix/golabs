package main

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func main() {
	var a, b, result int // var of type int

	fmt.Println("Enter number one: ")
	fmt.Scanln(&a)

	fmt.Println("Enter number two")
	fmt.Scanln(&b)

	result = sumar(a, b)

	fmt.Println("The result is: ", result)
}
