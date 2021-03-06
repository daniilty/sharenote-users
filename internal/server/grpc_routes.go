package server

import (
	"context"

	schema "github.com/daniilty/sharenote-grpc-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPC) AddUser(ctx context.Context, req *schema.AddUserRequest) (*schema.AddUserResponse, error) {
	uid, err := g.service.AddUser(ctx, convertPBAddUserToCore(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.AddUserResponse{
		Id: uid,
	}, nil
}

func (g *GRPC) GetUser(ctx context.Context, req *schema.GetUserRequest) (*schema.GetUserResponse, error) {
	user, ok, err := g.service.GetUser(ctx, req.GetId())
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetUserResponse{
		User: convertCoreUserToPB(user),
	}, nil
}

func (g *GRPC) GetUserByEmail(ctx context.Context, req *schema.GetUserByEmailRequest) (*schema.GetUserByEmailResponse, error) {
	user, ok, err := g.service.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetUserByEmailResponse{
		User: convertCoreUserToPB(user),
	}, nil
}

func (g *GRPC) IsValidUserCredentials(ctx context.Context, req *schema.IsValidUserCredentialsRequest) (*schema.IsValidUserCredentialsResponse, error) {
	isValid, err := g.service.IsValidUserCredentials(ctx, req.GetEmail(), req.GetPasswordHash())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.IsValidUserCredentialsResponse{
		IsValid: isValid,
	}, nil
}

func (g *GRPC) GetUsers(ctx context.Context, req *schema.GetUsersRequest) (*schema.GetUsersResponse, error) {
	users, err := g.service.GetUsers(ctx, req.GetIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetUsersResponse{
		Users: convertCoreUsersToPB(users),
	}, nil
}

func (g *GRPC) UpdateUser(ctx context.Context, req *schema.UpdateUserRequest) (*schema.UpdateUserResponse, error) {
	ok, err := g.service.UpdateUser(ctx, convertPBUpdateUserToCore(req))
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.UpdateUserResponse{}, nil
}
