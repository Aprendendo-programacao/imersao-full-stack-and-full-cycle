package data

import "github.com/Aprendendo-programacao/imersao-full-stack-and-full-cycle/entity"

var Products entity.Products

func LoadData() {
	product1 := entity.NewProduct()
	product1.Name = "Carrinho"

	product2 := entity.NewProduct()
	product2.Name = "Boneca"

	Products.Add(*product1)
	Products.Add(*product2)
}
