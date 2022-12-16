package main

import "fmt"

func saludar(name string) {
	fmt.Println("hola", name)
}

func main() {
	saludar("Tinix")
	developer()
	result := sumar(5, 3)
	fmt.Println("la suma es: ", result)
}

func developer() {
	fmt.Println("I'm Daniel Tinivella I'm Go Developer")
}

func sumar(a, b int) int {
	return a + b
}
