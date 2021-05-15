package main

import "fmt"

// Variables como referencias.

// Esta funcion recibe como parametro un puntero.
func cambiarValor(parametro *string) {
	*parametro = "Nuevo valor"
}

func main() {

	var variable string = "Valor"

	fmt.Printf("1) El valor de la variable es: %s | direccion de memoria: %p\n", variable, &variable)
	
	cambiarValor(&variable) // Se pasa la variable como una referencia.
	
	fmt.Printf("2) El valor de la variable es: %s | direccion de memoria: %p\n", variable, &variable)

}