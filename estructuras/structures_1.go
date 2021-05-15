package main

import "fmt"

// Definicion de estructuras.

type Usuario struct {
	nombre string
	edad int
}

func main(){

	// Crea una variable de tipo Usuario.
	var usuario Usuario

	usuario.nombre = "Fulano"
	usuario.edad = 25

	persona := Usuario{nombre: "Heraclito", edad: 24}

	var nombre string = "Descartes"
	var edad int = 50

	personaUsuario := Usuario{nombre, edad}

	fmt.Println(usuario)
	fmt.Println(persona)
	fmt.Println(personaUsuario)

}