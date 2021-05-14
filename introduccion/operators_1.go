package main

import "fmt"

// Operadores relacionales.
func main() {

	var a = 10
	var b = 5
	var c = 5
	var d = -1

	var txt_1 = "Texto"
	var txt_2 = "Texto"

	var esMayor = a > b
	var esMenor = b < a
	var esIgual = b == c
	var esDiferente = a != d

	fmt.Println(a, "es mayor que", b, "==", esMayor)
	fmt.Println(b, "es menor que", a, "==", esMenor)
	fmt.Println(b, "es igual que", c, "==", esIgual)
	fmt.Println(a, "es diferente que", d, "==", esDiferente)

	fmt.Println(txt_1, "es igual que", txt_2, "==", (txt_1 == txt_2))

	fmt.Println(a, "es igual que", b, "==", a == b)

}