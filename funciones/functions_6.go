package main

import "fmt"

// Funciones con N cantidad de argumentos (Variadic).

// La funcion promedio recibe N cantidad de numeros decimales, 
// por lo tanto se trata de una funcion Variadic.
func promedio(numeros ...float32)float32 {
	
	var suma float32

	for _, numero := range numeros {
		suma += numero
	}
	
	return suma / float32(len(numeros))
}

func main() {

	fmt.Println("El promedio es:", promedio(1.1, 2.2, 3.3))
	
}