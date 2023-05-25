package service

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository/order"
	ordermemory "github.com/KEMonster/ddd-go/domain/aggregate/repository/order/memory"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository/product"
	productmemory "github.com/KEMonster/ddd-go/domain/aggregate/repository/product/memory"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository/user"
	usermemory "github.com/KEMonster/ddd-go/domain/aggregate/repository/user/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	user    user.UserRepository
	product product.ProductRepository
	order   order.OrderRepository
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

func (os *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (uuid.UUID, error) {
	userinfo, err := os.user.Get(customerID)
	orderId, _ := uuid.NewUUID()
	if err != nil {
		return orderId, err
	}
	productInfo, err := os.product.GetAll(productIDs)
	if err != nil {
		return orderId, err
	}
	// Get each Product, Ouchie, We need a ProductRepository
	var price float64
	for _, p := range productInfo {
		price += p.GetPrice()
	}
	orders := entity.Order{
		OrderId:  orderId,
		User:     userinfo,
		Products: productInfo,
	}
	os.order.Add(orders.ToPO())
	return orderId, nil
}

// WithMemoryOrderRepository applies a memory order repository to the OrderService
func WithMemoryOrderRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	or := ordermemory.New()
	return WithOrderRepository(or)
}

// WithMemoryOrderRepository adds a in memory product repo and adds all input products
func WithOrderRepository(or order.OrderRepository) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		os.order = or
		return nil
	}
}

func WithMemoryUserRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	ur := usermemory.New()
	return WithUserRepository(ur)
}

func WithUserRepository(ur user.UserRepository) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		os.user = ur
		return nil
	}
}

func WithMemoryProductRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	pr := productmemory.New()
	return WithProductRepository(pr)
}

func WithProductRepository(pr product.ProductRepository) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		os.product = pr
		return nil
	}
}
