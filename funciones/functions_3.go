package main

import "fmt"

// Funciones como parametros.

type operacion func(a, b int) int // Creamos un nuevo tipo de dato.

func suma(a, b int) int { // Funcion tipo operacion.
	return a + b
}

func multiplicacion(a, b int) int { // Funcion tipo operacion.
	return a * b
}

func ejecutaOperacion(funcion operacion, a, b int){
	resultado := funcion(a, b)
	fmt.Println("Resultado:", resultado)
}

func main() {

	ejecutaOperacion(suma, 4, 3)
	ejecutaOperacion(multiplicacion, 3, 7)
	
}