package main

import "fmt"

func main() {

	var nombre string
	var edad int
	var altura float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s\n", &nombre) // Se pasa la variable como referencia.

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d\n", &edad)

	fmt.Print("Ingresa tu altura: ")
	fmt.Scanf("%f\n", &altura)

	fmt.Printf("\nTu nombre es %s tienes %d años y mides %.2f metros.\n", nombre, edad, altura)
}