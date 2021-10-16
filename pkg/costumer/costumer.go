package costumer

type Costumer struct {
	name    string
	address string
	phone   string
}

// New() returns a new customer
func New(name string, address string, phone string) Costumer {
	return Costumer{name, address, phone}
}
