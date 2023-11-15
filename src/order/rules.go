package order

type Rule interface {
	Satisfy(order Order) bool
}

type MinValueToDiscount10Percent struct {
	Value float64
}

func (minValue MinValueToDiscount10Percent) Satisfy(order Order) bool {
	return order.Total >= minValue.Value
}

type ExpectedCategory struct {
	Category string
}

func (expectCategory ExpectedCategory) Satisfy(o Order) bool {
	var hasExpectedCategory bool = false

	for _, product := range o.Products {
		hasExpectedCategory = product.Category == expectCategory.Category

		if hasExpectedCategory == true {
			break
		}
	}

	return hasExpectedCategory
}

type ExpectedPaymentMethod struct {
	PaymentMethod string
}

func (f ExpectedPaymentMethod) Satisfy(order Order) bool {
	return order.Payment.Method == f.PaymentMethod
}
