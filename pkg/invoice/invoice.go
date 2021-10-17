package invoice

import (
	costumer "github.com/ssgonzalezzz/POO-GO/pkg/costumer"
	"github.com/ssgonzalezzz/POO-GO/pkg/course"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  costumer.Costumer
	course  []course.Course
}

// New returns a new invoice
func New(country, city string, client costumer.Costumer, course []course.Course) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		course:  course,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.course {
		i.total += item.Price()
	}
}
