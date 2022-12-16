package main

import (
	"fmt"
)

func main() {
	// arrays
	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 10
	numbers[2] = 10
	numbers[3] = 10
	numbers[4] = 10

	fmt.Println(numbers)
	fmt.Println(numbers[4])

	// array with value
	names := [3]string{"Tinix", "George", "Tinivella"}

	fmt.Println("The array content: ", names)

	// definir un array sin longitud
	colors := [...]string{
		"red",
		"green",
		"black",
		"blue",
		"white",
	}

	fmt.Println(colors, len(colors))

	// Indices definidos
	monedas := [...]string{
		0: "Dolar",
		1: "Pesos",
		2: "Euros",
		3: "Rublo",
	}

	fmt.Println(monedas, len(monedas))

	subMonedas := monedas[0:3]
	fmt.Println(subMonedas)

	sub2Monedas := monedas[0:1]
	fmt.Println(sub2Monedas)

	last_monedas := monedas[3:]
	fmt.Println(last_monedas)
}
