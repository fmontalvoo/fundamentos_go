package main

import "fmt"

// Bloques y alcance de variables.
// Cada variable solo se puede acceder desde el bloque en el que fue declarada.
func main() { // Bloque #1

	var b1 = "Bloque#1"
	var version = 1
	fmt.Println("Bloque:", b1, "Version:", version)
	
	{ // Bloque #2
		
		var b2 = "Bloque#2"
		var version = 2
		fmt.Println("Bloque:", b1, "Version:", version)
		fmt.Println("Bloque:", b2, "Version:", version)
		
		{ // Bloque #3

			var b3 = "Bloque#3"
			var version = 3
			fmt.Println("Bloque:", b1, "Version:", version)
			fmt.Println("Bloque:", b2, "Version:", version)
			fmt.Println("Bloque:", b3, "Version:", version)
		
		}
		
		// fmt.Println("Bloque:", b3, "Version:", version)
	}
	
	// fmt.Println("Bloque:", b2, "Version:", version)
	// fmt.Println("Bloque:", b3, "Version:", version)
	
}