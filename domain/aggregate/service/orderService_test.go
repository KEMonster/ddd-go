package service

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/google/uuid"
	"testing"
)

func init_products(t *testing.T) []entity.Product {
	beer := entity.Product{uuid.New(), "Beer", 1.99}
	peenut := entity.Product{uuid.New(), "Peenut", 0.99}
	wine := entity.Product{uuid.New(), "Wine", 0.99}
	products := []entity.Product{
		beer, peenut, wine,
	}
	return products
}

func init_user(t *testing.T) entity.User {
	user := entity.User{uuid.New(), "KT", entity.Address{Country: "CN", Province: "Guangdong", Addr: "ShenZhen"}}
	return user
}

func TestOrderService_NewOrderService(t *testing.T) {
	products := init_products(t)
	user := init_user(t)

	os, err := NewOrderService(
		WithMemoryUserRepository(),
		WithMemoryProductRepository(),
		WithMemoryOrderRepository(),
	)
	os.user.Add(user)
	var productIds []uuid.UUID
	for _, product := range products {
		os.product.Add(product)
		productIds = append(productIds, product.GetId())
	}
	if err != nil {
		t.Error(err)
	}
	orderId, err := os.CreateOrder(user.GetUid(), productIds)
	if err != nil {
		t.Error(err)
	}
	t.Logf("orderId is %v", orderId)

}
