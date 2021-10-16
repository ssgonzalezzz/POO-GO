package main

import (
	"fmt"

	"github.com/ssgonzalezzz/POO-GO/course"
)

func main() {

	// Go.AddUserId(75)

	// fmt.Println(Go.UserIDs)

	// Go.AddClass("Prueba", "Prueba 2")

	// Go.PrintClasses()

	Go := course.New("Go desde cero", 15)
	Go.SetIds([]uint{12, 32, 84})
	Go.SetClasses(map[uint]string{
		1: "Introduccion",
		2: "Estructuras",
		3: "Mapas",
	})

	fmt.Printf("%+v", Go)

}
