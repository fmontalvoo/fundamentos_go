package main

import "fmt"

// Funciones recursivas.

func factorial(numero int)int {
	
	if numero == 1 {
		return 1
	}

	return numero * factorial(numero - 1)

}

func potencia(base, exponente int) int {
	if exponente == 0 {
		return 1
	}

	return base * potencia(base, exponente - 1)
}

func main() {
	
	fmt.Println("El factorial de 5 es:", factorial(5))

	// Se almacena la funcion en la variable.
	// var funcionPotencia = potencia

	// Declaracion de una variable de tipo funcion que recibe 2 argumentos y retorna un numero entero.
	var funcionPotencia func(b, e int) int
	funcionPotencia = potencia

	resultado := funcionPotencia(2, 10)

	fmt.Println("El numero 2 elevado a 8 es:", resultado)

}