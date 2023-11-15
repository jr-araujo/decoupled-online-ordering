package order

import (
	"github.com/google/uuid"
)

type orderRule struct {
	rule   Rule
	action Action
}

var rules []orderRule

type Order struct {
	OrderId        uuid.UUID
	Payment        Payment
	Products       []Product
	ShippingLabels []string
	Total          float64
}

func (o *Order) InitOrder() *Order {
	o.addRule(MinValueToDiscount10Percent{1000}, AddShippingLabel{"frete-gratis"})
	o.addRule(ExpectedCategory{"eletrodoméstico"}, AddShippingLabel{"frágil"})
	o.addRule(ExpectedPaymentMethod{"boleto"}, AddDiscountPercent{10})
	o.addRule(ExpectedCategory{"infantil"}, AddShippingLabel{"presente"})

	return &Order{}
}

func (o *Order) addRule(rule Rule, action Action) {
	rules = append(rules, struct {
		rule   Rule
		action Action
	}{rule, action})
}

func (o *Order) verifyRules(order Order) Order {
	for _, config := range rules {
		if config.rule.Satisfy(order) {
			order = config.action.Execute(order)
		}
	}
	return order
}

func (o *Order) AddProduct(product Product) {
	o.Products = append(o.Products, product)
	o.Total += float64(product.Value)
}

func (o *Order) PlaceOrder( /*payment Payment*/ ) Order {
	if len(o.Products) == 0 {
		panic("There isn't any product in the order!")
	}

	o.Payment.Value = float32(o.Total)

	var order Order

	for _, config := range rules {
		if config.rule.Satisfy(*o) {
			order = config.action.Execute(*o)
		}
	}

	return order
}
