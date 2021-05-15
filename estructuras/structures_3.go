package main

import "fmt"

// Estructuras embebidas.
// Relacion 1 a 1

type Persona struct {
	nombre string
	edad int
	activo bool
}

type Empleado struct {
	id string
	salario float32
	persona Persona
}

func main() {

	fulano := Persona{nombre: "Fulano", edad: 25, activo: true}
	fmt.Println("Persona:", fulano)

	fulanoEst := Empleado{id: "#01", salario:395.99, persona: fulano}
	fmt.Println("Empleado:", fulanoEst)

}