package main

import "fmt"

// Funcion panic.
func main() {

	var dividendo, divisor int

	fmt.Println("Ingresa el dividendo: ")
	fmt.Scanf("%d\n", &dividendo)

	fmt.Println("Ingresa el divisor: ")
	fmt.Scanf("%d\n", &divisor)

	if divisor == 0 {
		panic("No se puede dividir entre cero")
	}

	resultado := dividendo / divisor

	fmt.Println(dividendo, "/", divisor , "==", resultado)
	
}