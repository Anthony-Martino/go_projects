package account

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints ...
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

//MakeEndpoints ...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateUserRequest)
		if !ok {
			return nil, errors.New("Request is not of type CreateUserRequest")
		}
		return s.CreateUser(ctx, req.Email, req.Password)
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetUserRequest)
		if !ok {
			return nil, errors.New("Request is not of type GetUserRequest")
		}
		return s.GetUser(ctx, req.ID)
	}
}
