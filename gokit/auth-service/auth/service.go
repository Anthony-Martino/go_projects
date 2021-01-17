package auth

import (
	"context"

	"github.com/gofrs/uuid"
)

//Service ...
type Service interface {
	Login(ctx context.Context, req LoginRequest) (string, error)
	Register(ctx context.Context, req RegisterRequest) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

type service struct {
	repo Repository
}

//NewService ...
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s service) Login(ctx context.Context, req LoginRequest) (string, error) {
	user := User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.repo.Login(ctx, user)
	if err != nil {
		return "error logging in", err
	}
	return "Success", nil
}
func (s service) Register(ctx context.Context, req RegisterRequest) (string, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()

	user := User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.Register(ctx, user)
	if err != nil {
		return "failed to create user", err
	}
	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	email, err := s.repo.GetUser(ctx, id)

	if err != nil {
		return "", err
	}

	return email, nil
}
