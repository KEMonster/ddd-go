package entity

import (
	"github.com/google/uuid"
)

type OrderPo struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	ProductIds string
	Address    string
}

func (op *OrderPo) GetID() uuid.UUID {
	return op.Id
}
