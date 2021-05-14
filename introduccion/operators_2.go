package main

import "fmt"

// Operadores logicos.
func main() {

	 a := 10
	 b := 5
	 c := 5
	 d := -1

	 and := a > b && b > c
	 or := a != d || c > b
	 negation := !(a > b && (b != d || d > c))

	 fmt.Println("Operador And:", and)
	 fmt.Println("Operador Or:", or)
	 fmt.Println("Operador de negacion:", negation, !negation)

}