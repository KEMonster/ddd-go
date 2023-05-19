package service

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository"
	ordermemory "github.com/KEMonster/ddd-go/domain/aggregate/repository/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	user    repository.UserRepository
	product repository.ProductRepository
	order   repository.OrderRepository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderservice
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	userinfo, err := os.user.Get(customerID)
	if err != nil {
		return 0, err
	}
	productInfo, err := os.product.GetAll(productIDs)
	if err != nil {
		return 0, err
	}
	// Get each Product, Ouchie, We need a ProductRepository
	var price float64
	for _, p := range productInfo {
		price += p.GetPrice()
	}
	orderId, _ := uuid.NewUUID()
	orders := entity.Order{
		OrderId:  orderId,
		User:     userinfo,
		Products: productInfo,
	}
	os.order.Add(orders.ToPO())
	return price, nil
}

// WithMemoryOrderRepository applies a memory order repository to the OrderService
func WithMemoryOrderRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := ordermemory.New()
	return WithOrderRepository(cr)
}

// WithMemoryOrderRepository adds a in memory product repo and adds all input products
func WithOrderRepository(or repository.OrderRepository) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		os.order = or
		return nil
	}
}
