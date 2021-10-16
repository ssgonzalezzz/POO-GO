package course

import (
	"fmt"
)

// Estructura basica de los objetos de tipo Course. Pueden exportarse entre paquetes ya que su nombre inicia en mayusculas
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

//Imprime en consola las clases contenidas en un objeto de tipo Course.
func (c Course) PrintClasses() {

	if c.Classes != nil && len(c.Classes) != 0 {
		text := "Las clases del curso " + c.Name + " son: \n"

		for i := 1; i <= len(c.Classes); i++ {
			text = text + "- " + c.Classes[uint(i)] + "\n"
		}

		fmt.Println(text)
		return
	}
	fmt.Println("El curso no tiene clases asignadas")

}

// Permite agregar una o mas clases nuevas al curso
func (c *Course) AddClass(className ...string) {
	var l int
	for _, v := range className {
		l = len(c.Classes)
		c.Classes[uint(l)+1] = v
	}

}

//Agrega un nuevo usuario al curso
func (c *Course) AddUserId(id uint) []uint {
	c.UserIDs = append(c.UserIDs, id)
	return c.UserIDs
}
