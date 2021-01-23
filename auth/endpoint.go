package auth

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints ...
type Endpoints struct {
	Login    endpoint.Endpoint
	Register endpoint.Endpoint
	Token    endpoint.Endpoint
	GetUser  endpoint.Endpoint
}

//MakeEndpoints ...
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Login:    makeLoginEndpoint(s),
		Register: makeRegisterEndpoint(s),
		Token:    makeTokenEndpoint(s),
		GetUser:  makeGetUserEndpoint(s),
	}
}

func makeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(LoginRequest)
		if !ok {
			return nil, errors.New("Request is not of type LoginRequest")
		}

		return s.Login(ctx, req)
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

func makeTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(TokenRequest)
		if !ok {
			return nil, errors.New("Request is not of type TokenRequest")
		}

		return s.Token(ctx, req)
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
