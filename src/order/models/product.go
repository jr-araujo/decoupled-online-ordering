package order_models

import "github.com/google/uuid"

type Product struct {
	Id         uuid.UUID
	CategoryId string
	Name       string
	Value      int32
}

type Category struct {
	Id   string
	Name string
}
