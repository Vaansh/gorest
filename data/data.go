package data

import (
	"github.com/google/uuid"
)

type item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GenerateID() string {
	id := uuid.New()
	return id.String()
}

var Items = []item{
	{ID: GenerateID(), Name: "Item1", Price: 14.99},
	{ID: GenerateID(), Name: "Item2", Price: 24.00},
	{ID: GenerateID(), Name: "Item3", Price: 5.49},
}
