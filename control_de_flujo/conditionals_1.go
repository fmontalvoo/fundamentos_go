package main

import "fmt"

func main() {

	var esMayor bool
	var edad int

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d\n", &edad)

	esMayor = (edad >= 18)

	if esMayor {
		fmt.Println("Es mayor de edad")
	}else {
		fmt.Println("Es menor de edad")
	}

}