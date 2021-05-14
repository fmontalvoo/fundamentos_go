package main

import "fmt"

func main() {

	// Declaracion de variables con su tipo de dato.
	var nombre string
	var edad int

	nombre = "Frank"
	edad = 25

	// Declaracion de variables omitiendo su tipo de dato.
	apellido := "Montalvo"
	email := "fgmo@email.com"

	var altura = 1.75

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(email)
	fmt.Println(altura)
}