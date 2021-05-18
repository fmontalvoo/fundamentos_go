package main

import "fmt"

func main() {

	var a, b, c int = 6, 9, 3

	if a > b {
		if a > c {
			fmt.Println("A es el mayor")
		}else{
			fmt.Println("C es el mayor")
		}
	} else { 
		if b > c {
			fmt.Println("B es el mayor")
		} else {
			fmt.Println("C es el mayor")
		}
	}

	fmt.Println("+----------------+")

	if a > b && a > c {
		fmt.Println("A es el mayor")
	} else if b > a && b > c {
		fmt.Println("B es el mayor")
	} else {
		fmt.Println("C es el mayor")
	}

}