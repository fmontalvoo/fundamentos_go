package main

import "fmt"

// Declaracion de funciones.

func saluda(nombre string) {
	fmt.Println("Hola", nombre)
}

func suma(a, b int) int {
	return a + b
}

func sum(a, b int) (string, int) {
	return "Resultado suma:", a + b
}

func main() {
	saluda("Frank")
	fmt.Println("Suma:", suma(5, 2))
	
	texto, suma := sum(5, 2)

	fmt.Println(texto, suma)
}