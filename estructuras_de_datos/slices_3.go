package main

import "fmt"

// Funcion make
func main() {

	// Crea un slice de numeros enteros de longitud cero y con capacidad maxima de tres elementos.
	// slice := make([]int, 0, 3)

	slice := make([]int, 3, 3)

	fmt.Println(slice, len(slice), cap(slice))

	slice[0] = 3
	slice[1] = 5
	slice[2] = 7

	fmt.Println(slice)

	slice = append(slice, 11)

	fmt.Println(slice, len(slice), cap(slice))

}