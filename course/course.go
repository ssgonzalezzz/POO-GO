package course

import "fmt"

// Estructura basica de los objetos de tipo Course. Pueden exportarse entre paquetes ya que su nombre inicia en mayusculas
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

//Imprime en consola las clases contenidas en un objeto de tipo Course
func (c Course) PrintClasses() {
	text := "Las clases del curso " + c.Name + " son: \n"
	for _, v := range c.Classes {
		text += "- " + v + "\n"
	}
	fmt.Println(text)
}

//Agrega un nuevo usuario al curso
func (c *Course) AddUserId(id uint) []uint {
	c.UserIDs = append(c.UserIDs, id)
	return c.UserIDs
}
