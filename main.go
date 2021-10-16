package main

import (
	"fmt"

	"github.com/ssgonzalezzz/POO-GO/pkg/costumer"
	"github.com/ssgonzalezzz/POO-GO/pkg/invoice"
	"github.com/ssgonzalezzz/POO-GO/pkg/invoiceitem"
)

func main() {

	i := invoice.New(
		"Venezuela",
		"San Antonio de Los Altos",
		costumer.New("Samuel", "Res San Antonio", "+584125583846"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(2, "Curso de CSS", 10.00),
			invoiceitem.New(3, "Curso de JS", 25.15),
		},
	)

	i.SetTotal()

	fmt.Printf("%+v", i)

}
