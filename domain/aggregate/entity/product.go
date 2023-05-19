package entity

import "github.com/google/uuid"

type Product struct {
	Id    uuid.UUID
	Name  string
	Prize float64
}

func (p Product) GetPrice() float64 {
	return p.Prize
}

func (p Product) GetName() string {
	return p.Name
}

func (p Product) GetId() uuid.UUID {
	return p.Id
}
