package main

import "fmt"

func main() {
	
	// monedas := [...] string {"Sucres", "Pesos", "Soles"}

	// Indica de forma explicita el numero de indice donde se almacena un dato.
	monedas := [...] string {2:"Sucres", 1:"Pesos", 0:"Soles", 7:"Bolivares"}
	
	fmt.Println(monedas)

	// Sub arreglo del arreglo monedas.
	subArray := monedas[1:3]

	fmt.Println(subArray)

}