package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter of Item.value
func (i Item) Value() float64 { return i.value }
