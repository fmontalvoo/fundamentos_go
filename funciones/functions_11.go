package main

import "fmt"

// Ejecutar una funcion al finalizar la ejecucion el bloque donde es llamada.

func funcion1(){
	fmt.Println("Funcion #1")
}

func funcion2(){
	fmt.Println("Funcion #2")
}

func main() {

	fmt.Println("Funcion Main")

	// La palabra reservada `defer` obliga a la funcion1() a ejecutarse al finalizar el bloque de main().
	// Es similar a llamar la funcion al finla del bloque main().
	defer funcion1()
	funcion2()

}