package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Uid  uuid.UUID
	Name string
	Addr Address
}

func (u User) GetUid() uuid.UUID {
	return u.Uid
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAddress() Address {
	return u.Addr
}
