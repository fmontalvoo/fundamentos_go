package main

import "fmt"

// Declaracion de arreglos.
func main() {
	var numeros[5] int

	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3
	numeros[3] = 4
	numeros[4] = 5

	fmt.Println(numeros)

	// soremun := [5] int {1, 2, 3, 4, 5}
	
	// Si no se indica el numero de elementos, entonces go lo infiere.
	soremun := [...] int {1, 2, 3, 4, 5}

	fmt.Println(soremun)
}