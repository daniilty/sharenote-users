package core

import "github.com/daniilty/sharenote-users/internal/mongo"

func (n *User) toDB() *mongo.User {
	return &mongo.User{
		ID:             n.ID,
		Name:           n.Name,
		UserName:       n.UserName,
		Email:          n.Email,
		EmailConfirmed: n.EmailConfirmed,
		PasswordHash:   n.PasswordHash,
	}
}

func convertDBUserToService(user *mongo.User) *User {
	return &User{
		ID:             user.ID,
		Name:           user.Name,
		UserName:       user.UserName,
		Email:          user.Email,
		EmailConfirmed: user.EmailConfirmed,
		PasswordHash:   user.PasswordHash,
	}
}
