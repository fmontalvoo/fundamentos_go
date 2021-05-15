package main

import "fmt"

// Declaracion de slices.
func main() {

	numeros := [] int { 1, 2, 3, 4}

	// Agrega un nuevo elemento al final del arreglo.
	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)
	numeros = append(numeros, 8)

	fmt.Println(numeros)
	
	// Guarda una referencia al slice de numeros.
	nuevoSlice := numeros[0:4]

	numeros[0] = -1

	fmt.Println(numeros)

	fmt.Println(nuevoSlice)

}