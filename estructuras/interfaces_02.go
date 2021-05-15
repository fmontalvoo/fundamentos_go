package main

import "fmt"

// Implementacion de multiples interfaces.

type Animal interface{
	comer()
}

type Mascota interface{
	vacunas()
}

type Gato struct {
	nombre string
}

// Sobre escribe la funcion comer de la interface Animal para el Gato.
func (self Gato) comer() {
	fmt.Println(self.nombre, "come")
}

// Sobre escribe la funcion vacunas de la interface Mascota para el Gato.
func (self Gato) vacunas() {
	fmt.Println(self.nombre, "tiene todas sus vacunas")
}

func main() {

	var gato Gato = Gato{nombre: "Lucas"}
	gato.comer()
	gato.vacunas()

}