package paquete

// En go un metodo es publico cunado su nombre comienza con una letra Mayuscula.
// Los metodos privados son su opuesto, es decir, su nombre comienza con una letra minuscula.

// Metodos para la estructura de Pais.

// Metodo publico.
func (self *Pais) ObtenerPais()string {
	self.setId("01") // Asigna un id al pais/
	return self.id + ") Pais: " + self.getNombre()
}

// Metodo privado.
func (self *Pais) getNombre() string{
	return self.Nombre
}

// Metodo privado.
func (self *Pais) setId(id string) {
	self.id = id
}
