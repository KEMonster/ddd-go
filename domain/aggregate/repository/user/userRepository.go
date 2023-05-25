package user

import (
	"errors"
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("the user was not found in the repository")
	ErrUserAlreadyExist = errors.New("the user already exists")
)

type UserRepository interface {
	Get(uuid.UUID) (entity.User, error)
	Add(entity.User) error
}
