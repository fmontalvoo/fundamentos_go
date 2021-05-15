package main

import "fmt"

// Declaracion de variables.
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

	fmt.Println("Nombre:",nombre, "\nApellido:", apellido, "\nEdad:", edad, "\nEmail:", email, "\nAltura:", altura)
}