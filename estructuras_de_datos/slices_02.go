package main

import "fmt"

// Declaracion de slices.
func main() {

	var meses = [] string {"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", 
	"Agosto", "Septiembre"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v | la capacidad es: %v | direccion de memoria: %p\n", longitud, capacidad, meses)

	// Si el arreglo esta a su maxima capacidad entonces go crea uno nuevo.
	meses = append(meses, "Octubre")
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")
	
	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitud es: %v | la capacidad es: %v | direccion de memoria: %p\n", longitud, capacidad, meses)


}