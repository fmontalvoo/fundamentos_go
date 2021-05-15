package main

import "fmt"

// Declaracion de interfaces.

// Para implementar una interface se deben usar los mismos nombres de funciones,
// la misma catidad de parametros y tipo de retorno.

// Define la estructura y las funciones que realiza.
type Animal interface{
	comer()
	dormir()
}

type Gato struct {
	nombre string
}

// Sobre escribe la funcion comer para el Gato.
func (self Gato) comer() {
	fmt.Println(self.nombre, "come")
}

// Sobre escribe la funcion dormir para el Gato.
func (self Gato) dormir() {
	fmt.Println(self.nombre, "duerme")
}

// Ejecuta las funciones de los objetos que implementan la interface Animal.
func ejecutarFunciones(animal Animal){
	animal.comer()
	animal.dormir()
}

func main(){

	var gato = Gato{nombre: "Lucas"}
	ejecutarFunciones(gato)

}