package main

import "fmt"

// Bucle For implementado como un bucle While.
func main() {

	numero := 12345
	contador := 0

	fmt.Println("Numero:", numero)

	for numero > 0{
		numero /= 10
		contador++
	}
	
	fmt.Println("Cantidad de digitos:", contador)

}