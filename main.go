package main

import (
	"fmt"

	"github.com/ssgonzalezzz/POO-GO/pkg/costumer"
	"github.com/ssgonzalezzz/POO-GO/pkg/course"
	"github.com/ssgonzalezzz/POO-GO/pkg/invoice"
)

func main() {

	i := invoice.New(
		"Venezuela",
		"San Antonio de Los Altos",
		costumer.New("Samuel", "Res San Antonio", "+584125583846"),
		[]course.Course{
			course.New("Go desde cero", 7.8),
			course.New("POO con Go", 15.6),
			course.New("Testing con Go", 22.33),
		},
	)

	i.SetTotal()

	fmt.Printf("%+v", i)

}
