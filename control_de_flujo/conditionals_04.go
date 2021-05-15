package main

import "fmt"

func main() {

	usuarios := map[string] string{}

	usuarios["nombre"] = "Frank"
	usuarios["email"] = "fgmo@email.com"
	usuarios["valor"] = ""

	valor, ok := usuarios["valor"]

	if ok {
		fmt.Println("El valor obtenido es:", valor)
	} else {
		fmt.Println("No existe el valor")
	}

	if valor, ok = usuarios["llave"]; ok {
		fmt.Println("El valor obtenido es:", valor)
	} else {
		fmt.Println("No existe el valor")
	}

}