package main

// Comando para crear el modulo: go mod init modulos

import (
	"fmt"
	"modulos/paquete" // Importacion del modulo creado.
)

func main() {

	pais := paquete.Pais{Nombre: "Ecuador"}

	fmt.Println(pais.ObtenerPais())

}