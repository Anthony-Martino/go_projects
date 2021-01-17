package auth

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints ...
type Endpoints struct {
	Register endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

//MakeEndpoints ...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Register: makeRegisterEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(RegisterRequest)
		if !ok {
			return nil, errors.New("Request is not of type RegisterRequest")
		}

		return s.Register(ctx, req)
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
