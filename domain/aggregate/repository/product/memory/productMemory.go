package memory

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository/product"
	"github.com/google/uuid"
	"sync"
)

// ProductMemoryRepository fulfills ProductRepository interface
type ProductMemoryRepository struct {
	ProductMap map[uuid.UUID]entity.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of Order
func New() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		ProductMap: make(map[uuid.UUID]entity.Product),
	}
}

// Get a product from the repository
func (mr *ProductMemoryRepository) Get(productId uuid.UUID) (entity.Product, error) {
	if product, ok := mr.ProductMap[productId]; ok {
		return product, nil
	}
	return entity.Product{}, product.ErrProductNotFound
}

// GetAll product from the repository
func (mr *ProductMemoryRepository) GetAll(productIDs []uuid.UUID) ([]entity.Product, error) {
	// Collect all Products from map
	var products []entity.Product
	for _, productId := range productIDs {
		if product, ok := mr.ProductMap[productId]; ok {
			products = append(products, product)
		}
	}
	return products, nil
}

// Add will add a new product to the repository
func (mr *ProductMemoryRepository) Add(newprod entity.Product) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.ProductMap[newprod.GetId()]; ok {
		return product.ErrProductAlreadyExist
	}

	mr.ProductMap[newprod.GetId()] = newprod

	return nil
}
