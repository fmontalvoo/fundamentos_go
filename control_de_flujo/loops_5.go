package main

import "fmt"

// Bucle For implementado como Do While.
func main() {

	var numero int = 0

	for ok := true; ok; ok = (numero < 10) {
		fmt.Println(numero)
		numero++
	}

}