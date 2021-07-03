package model

type Product struct {
	ID   string
	Name string
}

func NewProduct() *Product {
	return &Product{}
}

type Products struct {
	Product []Product
}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) Add(product *Product) {
	p.Product = append(p.Product, *product)
}
