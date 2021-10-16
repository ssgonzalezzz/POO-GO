package course

import (
	"fmt"
	"strconv"
)

type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

//Metodo constructor del objeto
func New(name string, price float64) *course {
	var isFree bool = false

	if price == 0 {
		isFree = true
		fmt.Println("Precio = 0, se establecio como curso gratis")
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}

}

/* -------------------------------------------------------------------------- */
/*                              Setters y Getters                             */
/* -------------------------------------------------------------------------- */

// Setter para nombre de curso
func (c *course) SetName(name string) { c.name = name }

// Getter para nombre de curso
func (c *course) Name() string { return c.name }

// Setter para precio de curso
func (c *course) SetPrice(price float64) {
	if price == 0 {
		c.price = price
		c.isFree = true
		return
	}

	if c.isFree {
		c.isFree = false
		c.price = price
		return
	}
	c.price = price
}

//Getter para precio de curso
func (c *course) Price() float64 { return c.price }

//Getter para saber si el curso es gratuito. Para cambiar esta propiedad, usar el metodo SetPrice()
func (c *course) IsFree() bool { return c.isFree }

//Setter para IDs de usuarios
func (c *course) SetIds(ids []uint) { c.userIDs = ids }

// Setter clases de curso
func (c *course) SetClasses(classes map[uint]string) { c.classes = classes }

//Getter de classes de curso.
func (c *course) Classes() string {

	if c.classes != nil && len(c.classes) != 0 {
		text := "Las clases del curso " + c.name + " son: \n"

		for i, v := range c.classes {
			text += strconv.Itoa(int(i)) + "- " + v + "\n"
		}

		return text
	}
	return "El curso no tiene clases asignadas"

}

/* -------------------------------------------------------------------------- */
/*                            Fin Setters y Getters                           */
/* -------------------------------------------------------------------------- */

// Permite agregar una o mas clases nuevas al curso
func (c *course) AddClass(className ...string) {
	var l int
	for _, v := range className {
		l = len(c.classes)
		c.classes[uint(l)+1] = v
	}

}

//Agrega un nuevo usuario al curso
func (c *course) AddUserId(id uint) []uint {
	c.userIDs = append(c.userIDs, id)
	return c.userIDs
}
