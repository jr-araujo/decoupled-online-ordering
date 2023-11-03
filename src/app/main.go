package main

import (
	"github.com/google/uuid"
	order_models "jroberto.com/decoupled/online/ordering/order/models"
)

func main() {
	println("app running...!")

	order := new(order_models.Order)
	// order.PlaceOrder()
	order.AddProduct(order_models.Product{
		Id:         uuid.New(),
		CategoryId: "Test",
		Name:       "Product 1",
		Value:      0,
	})

	order.AddProduct(order_models.Product{
		Id:         uuid.New(),
		CategoryId: "Test",
		Name:       "Product 2",
		Value:      0,
	})
}
