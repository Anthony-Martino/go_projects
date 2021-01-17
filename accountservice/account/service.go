package account

import (
	"context"

	"github.com/gofrs/uuid"
)

//Service ...
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
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

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	err := s.repo.CreateUser(ctx, user)
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
