package main

import "fmt"

// Variables globales.

var username string // Variable global.

func funcion1(){
	username += "Mundo"
}

func funcion2(){
	username += "!"
}

func main() {

	username = "Hola "

	funcion1()
	funcion2()

	fmt.Println(username)

}