package main

import "fmt"

// Definicion de metodos para estructuras.

type Usuario struct {
	nombre string
	edad int
}

// Estos metodos solo podra ser utilizado para objetos
// de tipo Usuario.

// Esta funcion es un metodo de la estructura Usuario, 
// la cual permite asignar un nombre al mismo. 
func (self *Usuario) setNombre(nombre string){
	self.nombre = nombre 
}

func (self *Usuario) getNombre() string{
	return self.nombre
}

func (self *Usuario) setEdad(edad int){
	self.edad = edad
}

func (self *Usuario) getEdad() int{
	return self.edad
}

func main(){

	// Crea una variable de tipo Usuario.
	var usuario Usuario

	usuario.setNombre("Fulano")
	usuario.setEdad(21)

	fmt.Println("Nombre:", usuario.getNombre())
	fmt.Println("Edad:", usuario.getEdad())

}