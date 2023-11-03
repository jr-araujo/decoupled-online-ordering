package order_models

type Payment struct {
	TransactionId   string
	PaymentMethodId string
	Value           float32
}

type PaymentMethod struct {
	PaymentMethodId string
	Name            string
}
