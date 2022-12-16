package main

import "fmt"

func main() {
	h := "hola"
	m := "mundo!"

	fmt.Println(h)
	fmt.Println(m)

	name := "Tinix"
	edad := 48

	fmt.Printf("Hello, %s y su edad es %d \n", name, edad)
	fmt.Printf("Hello, %v y su edad es %d \n", name, edad)

	mensaje := fmt.Sprintf("hola, %s y su edad es %d", name, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", edad)

	fmt.Print("Ingrese otro nombre: ")

	fmt.Scanln(&name)

	fmt.Println("Otro nombre: ", name)
}
