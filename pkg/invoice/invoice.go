package invoice

import (
	costumer "github.com/ssgonzalezzz/POO-GO/pkg/costumer"
	"github.com/ssgonzalezzz/POO-GO/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  costumer.Costumer
	items   []invoiceitem.Item
}

// New returns a new invoice
func New(country, city string, client costumer.Costumer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
