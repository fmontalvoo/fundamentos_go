package main

import "fmt"

// Declaracion de mapas.
func main() {

	// MAP[KEY DATA TYPE]VALUE DATA TYPE
	dias := make(map[int]string)

	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"

	fmt.Println(dias)
	
	delete(dias, 0) // Elimina la llave 4 con su valor.
	
	fmt.Println(dias)
}