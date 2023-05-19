package repository

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	Get(productId uuid.UUID) (entity.Product, error)
	GetAll(productIds []uuid.UUID) ([]entity.Product, error)
}
