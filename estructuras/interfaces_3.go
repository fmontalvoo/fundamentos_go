package main

import "fmt"

// Uso de interfaces vacias.

// Funcion que puede recibir cualquier tipo de dato.
func mostrarVariables(objetos ...interface{}) {
	fmt.Printf("%v\n", objetos)
}

func main() {

	mostrarVariables(1, 2.2, true, 'A', "Hola")

}