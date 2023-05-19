package memory

import (
	"fmt"
	"sync"

	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository"
	"github.com/google/uuid"
)

// OrderMemoryRepository fulfills OrderRepository interface
type OrderMemoryRepository struct {
	Order map[uuid.UUID]entity.OrderPo
	sync.Mutex
}

// New is a factory function to generate a new repository of Order
func New() *OrderMemoryRepository {
	return &OrderMemoryRepository{
		Order: make(map[uuid.UUID]entity.OrderPo),
	}
}

// Add will add a new customer to the repository
func (mr *OrderMemoryRepository) Add(c entity.OrderPo) error {
	if mr.Order == nil {
		// Saftey check if Order is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.Order = make(map[uuid.UUID]entity.OrderPo)
		mr.Unlock()
	}
	// Make sure Order isn't already in the repository
	if _, ok := mr.Order[c.GetID()]; ok {
		return fmt.Errorf("order already exists: %w", repository.ErrFailedToAddOrder)
	}
	mr.Lock()
	mr.Order[c.GetID()] = c
	mr.Unlock()
	return nil
}
