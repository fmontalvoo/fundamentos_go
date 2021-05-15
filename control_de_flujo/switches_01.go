package main

import "fmt"

func main(){

	a, b, c := 6, 9, 3

	switch{
	case a > b && b > c:
		fmt.Println("A es el mayor")
	case b > a && a > c:
		fmt.Println("B es el mayor")
	default:
		fmt.Println("C es el mayor")
	
	}

	var calificacion int
	fmt.Println("Ingrese su calificacion: ")
	fmt.Scanf("%d\n", &calificacion)

	switch calificacion {
	case 10:
		fmt.Println("Sobresaliente")
	case 8, 9:
		fmt.Println("Muy Buena")
	case 6, 7:
		fmt.Println("Buena")
	case 5:
		fmt.Println("Regular")
	case 1, 2, 3, 4:
		fmt.Println("Pesimo")
	default:
		fmt.Println("No valido")
	}

}