package server

import (
	schema "github.com/daniilty/sharenote-grpc-schema"
	"github.com/daniilty/sharenote-users/internal/core"
)

type GRPC struct {
	schema.UnimplementedUsersServer

	service core.Service
}

func NewGRPC(usersService core.Service) *GRPC {
	return &GRPC{
		service: usersService,
	}
}
