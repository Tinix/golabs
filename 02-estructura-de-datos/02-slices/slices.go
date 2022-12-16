package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3}
	fmt.Println(numeros)

	fmt.Println(numeros, len(numeros))

	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros)
	fmt.Println("Numeros length", len(numeros))

	subSlices := numeros[:2]

	numeros[0] = 100
	fmt.Println(subSlices)
	fmt.Println(numeros)

	// caracteristicas de los slices
	// Punteros
	// Longitud
	// Capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Length: %v, Capacity: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf("Length: %v, Capacity: %v, %p \n", len(meses), cap(meses), meses)
}
