package order

import "github.com/google/uuid"

type Product struct {
	Id       uuid.UUID
	Category string
	Name     string
	Value    int32
}
