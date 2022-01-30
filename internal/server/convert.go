package server

import (
	schema "github.com/daniilty/sharenote-grpc-schema"
	"github.com/daniilty/sharenote-users/internal/core"
)

func convertPBAddUserToCore(user *schema.AddUserRequest) *core.User {
	return &core.User{
		Name:         user.Name,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func convertPBUpdateUserToCore(user *schema.UpdateUserRequest) *core.User {
	return &core.User{
		Name:         user.Name,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func convertCoreUserToPB(user *core.User) *schema.User {
	return &schema.User{
		Id:             user.ID,
		Name:           user.Name,
		UserName:       user.UserName,
		Email:          user.Email,
		EmailConfirmed: user.EmailConfirmed,
	}
}

func convertCoreUserToResponse(u *core.User) user {
	return user{
		ID:       u.ID,
		Name:     u.Name,
		UserName: u.UserName,
	}
}

func convertPBUserToCore(user *schema.User) *core.User {
	return &core.User{
		ID:             user.Id,
		Name:           user.Name,
		UserName:       user.UserName,
		Email:          user.Email,
		EmailConfirmed: user.EmailConfirmed,
	}
}

func convertCoreUsersToPB(users []*core.User) []*schema.User {
	converted := make([]*schema.User, 0, len(users))

	for i := range users {
		converted = append(converted, convertCoreUserToPB(users[i]))
	}

	return converted
}

func convertCoreUsersToResponse(users []*core.User) []user {
	converted := make([]user, 0, len(users))

	for i := range users {
		converted = append(converted, convertCoreUserToResponse(users[i]))
	}

	return converted
}
