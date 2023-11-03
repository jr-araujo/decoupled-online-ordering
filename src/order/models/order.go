package order_models

import "github.com/google/uuid"

type Order struct {
	OrderId       uuid.UUID
	TransactionId uuid.UUID
	Products      []Product
	Items         []LineItem
	Labels        []string
	Total         float64
	// OrderPlacedAt
}

type LineItem struct {
	Id        int32
	ProductId uuid.UUID
	Quantity  int16
	Value     float64
}

// TODO: Before adding a product, search by using a ProductId in a product list and then add it into the o.Items
func (o *Order) AddProduct(product Product) {
	o.Products = append(o.Products, product)
}

func (o *Order) PlaceOrder() {
	if len(o.Products) == 0 {
		panic("There isn't any product in the order!")
	}

	for _, product := range o.Products {
		if product.Value > 1000 {
			o.Labels = append(o.Labels, "frete grátis")
		}

		// TODO: search CategoryId in a category list
		if product.CategoryId == "eletrodoméstico" {
			o.Labels = append(o.Labels, "frágil")
		}

		// TODO: search CategoryId in a category list
		if product.CategoryId == "infantil" {
			o.Labels = append(o.Labels, "presente")
		}

		o.Total += float64(product.Value)
	}

	// TODO: Search what is the payment method of this TransactionId
	// if o.TransactionId == "boleto" {
	// 	o.Total = o.Total * 0.9
	// }

}
