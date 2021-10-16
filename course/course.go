package course

import (
	"fmt"
	"strconv"
)

// Estructura basica de los objetos de tipo Course. Pueden exportarse entre paquetes ya que su nombre inicia en mayusculas
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

//Metodo constructor del objeto
func New(name string, price float64) *course {
	var isFree bool = false

	if price == 0 {
		isFree = true
		fmt.Println("Precio = 0, se establecio como curso gratis")
	}

	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}

}

//Imprime en consola las clases contenidas en un objeto de tipo Course.
func (c *course) PrintClasses() {

	if c.Classes != nil && len(c.Classes) != 0 {
		text := "Las clases del curso " + c.Name + " son: \n"

		for i, v := range c.Classes {
			text += strconv.Itoa(int(i)) + "- " + v + "\n"
		}

		fmt.Println(text)
		return
	}
	fmt.Println("El curso no tiene clases asignadas")

}

// Permite agregar una o mas clases nuevas al curso
func (c *course) AddClass(className ...string) {
	var l int
	for _, v := range className {
		l = len(c.Classes)
		c.Classes[uint(l)+1] = v
	}

}

//Agrega un nuevo usuario al curso
func (c *course) AddUserId(id uint) []uint {
	c.UserIDs = append(c.UserIDs, id)
	return c.UserIDs
}
