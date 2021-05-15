package main

import "fmt"

// Declaracion de variables dentro de la condicion.
func main() {

	// Las variables declaradas en el bloque if solo funcionaran a dentro de este.
	if nombre, edad := "Fulano", 17; nombre == "Fulano" && edad >= 18{
		fmt.Println("Hola", nombre)
	}else{
		fmt.Println(nombre, edad)
	}

}