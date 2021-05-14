package main

import "fmt"
import "reflect"

func main() {

	// var texto string = "Este es un texto."
	// var texto = "Este es un texto."
	texto := "Este es un texto."

	// Obtener la cantidad de caracteres.
	length := len(texto)

	fmt.Println("Cantidad de caracteres:", length)

	// Obteniendo un caracter en una posicion dentro del texto.
	primerCaracter := texto[0] // Char -> uint8
	ultimoCaracter := texto[length-1]

	fmt.Println("Codigo ascii del primer caracter:", primerCaracter)
	fmt.Println(reflect.TypeOf(primerCaracter)) // Muestra el tipo de dato.
	
	fmt.Printf("%c\n", primerCaracter)
	fmt.Printf("%c\n", ultimoCaracter)

}