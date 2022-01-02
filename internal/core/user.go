package core

import (
	"context"
)

type User struct {
	ID             string
	Name           string
	UserName       string
	Email          string
	EmailConfirmed bool
	PasswordHash   string
}

func (s *ServiceImpl) AddUser(ctx context.Context, user *User) error {
	return s.db.AddUser(ctx, user.toDB())
}

func (s *ServiceImpl) GetUser(ctx context.Context, id string) (*User, bool, error) {
	user, ok, err := s.db.GetUser(ctx, id)
	if err != nil {
		return nil, ok, err
	}

	return convertDBUserToService(user), true, nil
}

func (s *ServiceImpl) GetUsers(ctx context.Context, ids []string) ([]*User, error) {
	users := make([]*User, 0, len(ids))

	for i := range ids {
		user, ok, err := s.db.GetUser(ctx, ids[i])
		if err != nil {
			if ok {
				continue
			}

			return nil, err
		}

		users = append(users, convertDBUserToService(user))
	}

	return users, nil
}

func (s *ServiceImpl) UpdateUser(ctx context.Context, user *User) (bool, error) {
	return s.db.UpdateUser(ctx, user.toDB())
}
