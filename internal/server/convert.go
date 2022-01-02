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

func convertPBUserToCore(user *schema.User) *core.User {
	return &core.User{
		ID:             user.Id,
		Name:           user.Name,
		UserName:       user.UserName,
		Email:          user.Email,
		EmailConfirmed: user.EmailConfirmed,
	}
}

func convertCoreUsersToPB(notes []*core.User) []*schema.User {
	converted := make([]*schema.User, 0, len(notes))

	for i := range notes {
		converted = append(converted, convertCoreUserToPB(notes[i]))
	}

	return converted
}
