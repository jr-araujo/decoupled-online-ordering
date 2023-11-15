package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPlaceOrder_With_OneFreeShippingItem(t *testing.T) {
	order := new(Order)
	order = order.InitOrder()
	order.Payment = Payment{Method: "a vista"}

	order.AddProduct(Product{
		Id:       uuid.New(),
		Category: "casa",
		Name:     "Product 1",
		Value:    1001,
	})

	order.AddProduct(Product{
		Id:       uuid.New(),
		Category: "moda masculina",
		Name:     "Product 2",
		Value:    100,
	})

	var placedOrder = order.PlaceOrder()

	assert.Equal(t, 1, len(placedOrder.ShippingLabels))
	assert.Equal(t, "frete-gratis", placedOrder.ShippingLabels[0])
	assert.Equal(t, float64(1101), placedOrder.Total)
}

func TestPlaceOrder_With_OneGiftItem(t *testing.T) {
	order := new(Order)
	order = order.InitOrder()
	order.Payment = Payment{Method: "a vista"}

	order.AddProduct(Product{
		Id:       uuid.New(),
		Category: "infantil",
		Name:     "Product 3",
		Value:    100,
	})

	var placedOrder = order.PlaceOrder()

	assert.Equal(t, 1, len(placedOrder.ShippingLabels))
	assert.Equal(t, "presente", placedOrder.ShippingLabels[0])
	assert.Equal(t, float64(100), placedOrder.Total)
}

func TestPlaceOrder_With_OneFragileItem(t *testing.T) {
	order := new(Order)
	order = order.InitOrder()
	order.Payment = Payment{Method: "a vista"}

	order.AddProduct(Product{
		Id:       uuid.New(),
		Category: "eletrodoméstico",
		Name:     "Product 4",
		Value:    300,
	})

	var placedOrder = order.PlaceOrder()

	assert.Equal(t, 1, len(placedOrder.ShippingLabels))
	assert.Equal(t, "frágil", placedOrder.ShippingLabels[0])
	assert.Equal(t, float64(300), placedOrder.Total)
}

func TestPlaceOrder_WithDiscount10Percent(t *testing.T) {
	order := new(Order)
	order = order.InitOrder()
	order.Payment = Payment{Method: "boleto"}

	order.AddProduct(Product{
		Id:       uuid.New(),
		Category: "moda masculina",
		Name:     "Product 5",
		Value:    1500,
	})

	var placedOrder = order.PlaceOrder()

	assert.Equal(t, float32(1350), placedOrder.Payment.Value)
}

func TestPlaceOrder_WithoutProducts(t *testing.T) {
	order := new(Order)
	order = order.InitOrder()
	order.Payment = Payment{Method: "boleto"}

	assert.Panics(t, func() { order.PlaceOrder() })
}
