package main

import "fmt"

// Estructuras embebidas.
// Relacion 1 a *

type Persona struct {
	id int
	nombre string
	telefonos []Telefono
}

type Telefono struct {
	numero string
	persona Persona
}

func main() {

	telefono1 := Telefono{numero: "0987654321"}
	telefono2 := Telefono{numero: "0987456321"}

	telefonos := []Telefono { telefono1, telefono2 }

	fulano := Persona{id: 1, nombre: "Fulano", telefonos: telefonos}

	telefono1.persona = fulano
	telefono2.persona = fulano

	fmt.Println("Telefono:", telefono1)
	fmt.Println("Telefono:", telefono2)

	fmt.Println("Persona:", fulano)

	fmt.Println(fulano.nombre, fulano.telefonos[0].numero)
	fmt.Println(telefono2.numero, telefono2.persona.nombre)

	for _, telefono := range fulano.telefonos {
		fmt.Println(telefono.numero)
	}

}