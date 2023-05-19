package repository

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Get(uuid uuid.UUID) (entity.User, error)
}
