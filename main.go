package main

import (
	"fmt"

	"github.com/ssgonzalezzz/POO-GO/course"
)

func main() {

	Go := course.Course{
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

	Go.AddUserId(75)

	fmt.Println(Go.UserIDs)

	Go.AddClass("Prueba")

	Go.AddClass("Prueba 2")

	Go.PrintClasses()

}
