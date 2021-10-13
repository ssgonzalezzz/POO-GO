package main

import (
	"fmt"
)

type Curso struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c Curso) PrintClasses() {
	text := "Las clases del curso " + c.Name + " son: \n"
	for _, v := range c.Classes {
		text += "- " + v + "\n"
	}
	fmt.Println(text)
}

func (c *Curso) AddUserId(id uint) []uint {
	c.UserIDs = append(c.UserIDs, id)
	return c.UserIDs
}

func main() {

	Go := Curso{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 32, 84},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Mapas",
		},
	}

	(&Go).AddUserId(69)
	fmt.Println(Go.UserIDs)

	// Css := Curso{
	// Name:  "CSS desde cero",
	// Price: 20,
	// }
	//
	// Js := Curso{}
	//
	// Js.Name = "Javascript desde cero"

	// fmt.Println(Go.Name)
	// fmt.Printf("%+v\n", Css)
	// fmt.Println(Js.Name)

}
