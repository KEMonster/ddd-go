package memory

import (
	"github.com/KEMonster/ddd-go/domain/aggregate/entity"
	"github.com/KEMonster/ddd-go/domain/aggregate/repository/user"
	"github.com/google/uuid"
	"sync"
)

// UserMemoryRepository fulfills UserRepository interface
type UserMemoryRepository struct {
	UserMap map[uuid.UUID]entity.User
	sync.Mutex
}

// New is a factory function to generate a new repository of Order
func New() *UserMemoryRepository {
	return &UserMemoryRepository{
		UserMap: make(map[uuid.UUID]entity.User),
	}
}

// Get a user from the repository
func (mr *UserMemoryRepository) Get(userId uuid.UUID) (entity.User, error) {
	if user, ok := mr.UserMap[userId]; ok {
		return user, nil
	}
	return entity.User{}, user.ErrUserNotFound
}

// Add will add a new user to the repository
func (mr *UserMemoryRepository) Add(newuser entity.User) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.UserMap[newuser.GetUid()]; ok {
		return user.ErrUserAlreadyExist
	}

	mr.UserMap[newuser.GetUid()] = newuser

	return nil
}
