package order

import (
	"errors"
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
)

var (
	// ErrOrderNotFound is returned when a order is not found.
	ErrOrderNotFound = errors.New("the order was not found in the repository")
	// ErrFailedToAddOrder is returned when the order could not be added to the repository.
	ErrFailedToAddOrder = errors.New("failed to add the order to the repository")
	// ErrUpdateOrder is returned when the order could not be updated in the repository.
	ErrUpdateOrder = errors.New("failed to update the order in the repository")
)

type OrderRepository interface {
	Add(entity.OrderPo) error
}
