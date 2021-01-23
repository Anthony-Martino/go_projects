package auth

import "context"

//Service ...
type Service interface {
	Login(ctx context.Context, req LoginRequest) (string, error)
	Register(ctx context.Context, req RegisterRequest) (string, error)
	Token(ctx context.Context, req TokenRequest) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
