package main

import (
	"fmt"
	"errors"
)

// Manejo de errores en funciones.

// Esta funcion retorna un numero decimal y un error.
func division(dividendo float32, divisor float32) (float32, error) {
	if divisor == 0 {
		return -1, errors.New("No es posible dividir entre cero.")
	}
	return (dividendo / divisor), nil
}


func main() {

	if division1, err := division(37, 7); err != nil {
		fmt.Println("Error:", err)
	}else {
		fmt.Println("37 / 7 =", division1)
	}
	
	if division2, err := division(7, 0); err != nil {
		panic(err)
	}else {
		fmt.Println("7 / 0 =", division2)
	}

}