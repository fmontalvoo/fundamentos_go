package main

import "fmt"

// Declaracion de funciones anonimas.
func main() {

	// Define y ejecuta la funcion.
	func (){
		fmt.Println("Funcion anonima")
	}()

	// Almacena la funcion en una variable.
	funcionSuma := func (a,b int) int{
		return a + b
	}

	fmt.Println(funcionSuma(2, 5))
	
}