package main

import "fmt"

// Bucle Foreach.
func main() {

	var animales = [...]string {"Gato", "Perro", "Hamster"}

	for indice, elemento := range animales {
		fmt.Println(indice, elemento)
	}

	for _, animal := range animales {
		fmt.Println(animal)
	}

}