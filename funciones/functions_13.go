package main

import (
	"fmt"
	"os"
)

// Funcion recover.

func main() {

	terminarEjecucion := func(){
		// La funcion recover() retorna el error que lanzo la funcion panic().
		if err := recover(); err != nil {
			fmt.Println("Parece que algo salio mal :(")
		}
	}

	// Esta funcion permite retomar el control del programa 
	// y terminar su ejecucion sin un estado de error.
	defer terminarEjecucion()

	if file, err := os.Open("files.txt"); err == nil {
		// Esta funcion anonima se ejecutara al finalizar la ejecucion de la funcion principal (main).
		defer func() {
			file.Close() // Cierra el archivo file.txt
		}()

		contenido := make([]byte, 254)

		longitud, _ := file.Read(contenido)

		bytes := contenido[0:longitud] // Recupera el contenido como un slice de bytes.
		fmt.Println("Contenido en bytes:", bytes)
		
		texto := string(bytes) // Convierte el slice de bytes a texto.
		fmt.Println("Contenido en texto:\n", texto)

	} else {
		panic("No se puede acceder al archivo")
	}

}