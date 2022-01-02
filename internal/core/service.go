package core

import (
	"context"

	"github.com/daniilty/sharenote-users/internal/mongo"
)

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	// AddUser - add user to database.
	AddUser(context.Context, *User) error
	// GetNote - get user by id.
	GetUser(context.Context, string) (*User, bool, error)
	// GetUsers - get users by id.
	GetUsers(context.Context, []string) ([]*User, error)
	// UpdateUser - update user by id.
	UpdateUser(context.Context, *User) (bool, error)
}

type ServiceImpl struct {
	db mongo.DB
}

func NewServiceImpl(db mongo.DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
