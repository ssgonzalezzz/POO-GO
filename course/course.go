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

// Permite agregar una nueva clase al curso
func (c *Course) AddClass(className string) {
	l := len(c.Classes)
	c.Classes[uint(l)+1] = className
}

//Agrega un nuevo usuario al curso
func (c *Course) AddUserId(id uint) []uint {
	c.UserIDs = append(c.UserIDs, id)
	return c.UserIDs
}
