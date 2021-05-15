package main

import "fmt"

// Retorno de funciones.

type operacion func(a, b int) int // Creamos un nuevo tipo de dato.

func creaOperacion(tipo string)operacion{
	switch tipo {
	case "suma":
		return func(a, b int) int {
			return a + b
		}
	case "multiplicacion":
		return func(a, b int) int {
			return a * b
		}
	default:
		return func(a, b int) int {
			return a - b
		}
	}
}

func main() {

	funcionSuma := creaOperacion("suma")
	fmt.Println("Suma:", funcionSuma(6, 1))
	
	funcionMulticacion := creaOperacion("multiplicacion")
	fmt.Println("Multicacion:", funcionMulticacion(7, 4))

	funcionResta := creaOperacion("")
	fmt.Println("Resta:", funcionResta(6, 2))
	
}