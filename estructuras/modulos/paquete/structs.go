package paquete

// Los atributos de una estructura son publicos cuando comienzan por una letra Mayuscula,
// Y son privados cuando comienzan por una letra minuscula.

type Pais struct {
	id string // privado 
	Nombre string // publico
}