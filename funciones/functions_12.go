package main

import (
	"fmt"
	"os"
)

// Ejecutar una funcion al finalizar la ejecucion el bloque donde es llamada.

func main() {

	fmt.Println("Abre el archivo")
	file, err := os.Open("file.txt") // Abre el archivo file.txt
	
	if err != nil {
		panic(err)
	}

	// Arreglo de bytes de longitud 254.
	contenido := make([]byte, 254)

	longitud, _ := file.Read(contenido)

	bytes := contenido[0:longitud] // Recupera el contenido como un slice de bytes.
	fmt.Println("Contenido en bytes:", bytes)
	
	texto := string(bytes) // Convierte el slice de bytes a texto.
	fmt.Println("Contenido en texto:\n", texto)

	// Esta funcion anonima se ejecutara al finalizar la ejecucion de la funcion principal (main).
	defer func() {
		fmt.Println("Cierra el archivo")
		file.Close() // Cierra el archivo file.txt
	}()

}