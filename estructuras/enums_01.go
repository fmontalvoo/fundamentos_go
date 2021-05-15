package main

import "fmt"

// Declaracion de enums.

type TipoUsuario int

const (
	Cliente TipoUsuario = 1
	Administrador TipoUsuario = 2
)

type Usuario struct {
	nombre string
	tipo TipoUsuario
}

func main() {

	fulano := Usuario{nombre: "Fulano", tipo: Administrador}
	fulanito := Usuario{nombre: "Fulanito", tipo: Cliente}

	fmt.Println("Administrador:", fulano)
	fmt.Println("Cliente:", fulanito)

}