package main

import "fmt"

// Declaracion de constantes en secuencia.

// Declaracion de multiples constantes.
const (
	// Inicia con el valor cero y asigna valores en secuencia a las demas constantes.
	// Para iniciar con un valor distinto de cero podemos sumar `iota + 1` por ejemplo
	// entonces la secuencia inicia desde 1 hasta 7.
	Domingo int = iota 
	Lunes // 1
	Martes // 2
	Miercoles // 3
	Jueves // 4
	Viernes // 5
	Sabado // 6
)

func main() {
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}