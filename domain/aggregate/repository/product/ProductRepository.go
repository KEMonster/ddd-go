package product

import (
	"errors"
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("the product was not found in the repository")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	Get(uuid.UUID) (entity.Product, error)
	GetAll([]uuid.UUID) ([]entity.Product, error)
	Add(entity.Product) error
}
